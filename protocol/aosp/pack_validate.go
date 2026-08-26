package aosp

import (
	"crypto/sha256"
	"encoding/binary"
	"os"
	"path/filepath"
	"strings"
	"sync"

	corepack "github.com/i-rtfsc/as-aosp/core/pack"
	"github.com/i-rtfsc/as-aosp/protocol/jsonrpc"
)

type packValidationStamp struct {
	rootTime, manifestTime, checksumsTime int64
	manifestSize, checksumsSize           int64
	filesDigest                           [32]byte
}

type packValidationSnapshot struct {
	stamp  packValidationStamp
	result corepack.Result
}

type packValidationFlight struct {
	done   chan struct{}
	stamp  packValidationStamp
	result corepack.Result
}

type packValidationCache struct {
	mu      sync.Mutex
	entries map[string]packValidationSnapshot
	flights map[string]*packValidationFlight
}

func (server *Server) validatePackIfDeclared(root string) []corepack.Diagnostic {
	stamp, declared := currentPackStamp(root)
	if !declared {
		return nil
	}
	for {
		server.packValidation.mu.Lock()
		if snapshot, ok := server.packValidation.entries[root]; ok && snapshot.stamp == stamp {
			server.packValidation.mu.Unlock()
			return snapshot.result.Diagnostics
		}
		if flight, ok := server.packValidation.flights[root]; ok && flight.stamp == stamp {
			server.packValidation.mu.Unlock()
			<-flight.done
			stamp, declared = currentPackStamp(root)
			if !declared {
				return nil
			}
			continue
		}
		if server.packValidation.flights == nil {
			server.packValidation.flights = make(map[string]*packValidationFlight)
		}
		flight := &packValidationFlight{done: make(chan struct{}), stamp: stamp}
		server.packValidation.flights[root] = flight
		server.packValidation.mu.Unlock()
		result := corepack.Validate(root)
		endStamp, endDeclared := currentPackStamp(root)
		server.packValidation.mu.Lock()
		delete(server.packValidation.flights, root)
		if endDeclared && endStamp == stamp {
			if server.packValidation.entries == nil {
				server.packValidation.entries = make(map[string]packValidationSnapshot)
			}
			server.packValidation.entries[root] = packValidationSnapshot{stamp: stamp, result: result}
		}
		flight.result = result
		close(flight.done)
		server.packValidation.mu.Unlock()
		if endDeclared && endStamp != stamp {
			stamp = endStamp
			continue
		}
		return result.Diagnostics
	}
}

func currentPackStamp(root string) (packValidationStamp, bool) {
	rootInfo, rootErr := os.Stat(root)
	manifestInfo, manifestErr := os.Stat(filepath.Join(root, "manifest.json"))
	checksumsInfo, checksumsErr := os.Stat(filepath.Join(root, "checksums.sha256"))
	if os.IsNotExist(manifestErr) {
		return packValidationStamp{}, false
	}
	if rootErr != nil || manifestErr != nil || checksumsErr != nil {
		return packValidationStamp{}, true
	}
	digest, err := packFilesDigest(root)
	if err != nil {
		return packValidationStamp{}, true
	}
	return packValidationStamp{
		rootTime: rootInfo.ModTime().UnixNano(), manifestTime: manifestInfo.ModTime().UnixNano(),
		checksumsTime: checksumsInfo.ModTime().UnixNano(), manifestSize: manifestInfo.Size(),
		checksumsSize: checksumsInfo.Size(),
		filesDigest:   digest,
	}, true
}

func packFilesDigest(root string) ([32]byte, error) {
	hash := sha256.New()
	err := filepath.WalkDir(root, func(pathName string, entry os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() || entry.Name() == "checksums.sha256" {
			return nil
		}
		relative, err := filepath.Rel(root, pathName)
		if err != nil {
			return err
		}
		info, err := entry.Info()
		if err != nil {
			return err
		}
		_, _ = hash.Write([]byte(filepath.ToSlash(relative)))
		var buffer [16]byte
		binary.LittleEndian.PutUint64(buffer[:8], uint64(info.Size()))
		binary.LittleEndian.PutUint64(buffer[8:], uint64(info.ModTime().UnixNano()))
		_, _ = hash.Write(buffer[:])
		return nil
	})
	if err != nil {
		return [32]byte{}, err
	}
	return sha256.Sum256(hash.Sum(nil)), nil
}

type PackValidateParams struct {
	MetadataPackRoot string `json:"metadataPackRoot"`
}

type PackValidateResult struct {
	Valid       bool                  `json:"valid"`
	Diagnostics []corepack.Diagnostic `json:"diagnostics"`
}

func (server *Server) packValidate(request jsonrpc.Request) jsonrpc.Response {
	var params PackValidateParams
	if err := decodeParams(request, &params); err != nil {
		return jsonrpc.Failure(request.ID, jsonrpc.InvalidParams(err.Error()))
	}
	root := strings.TrimSpace(params.MetadataPackRoot)
	if root == "" {
		return jsonrpc.Success(request.ID, PackValidateResult{Diagnostics: []corepack.Diagnostic{{Code: "PACK_ROOT_MISSING", Message: "metadataPackRoot is required"}}})
	}
	if _, err := os.Stat(filepath.Join(root, "manifest.json")); os.IsNotExist(err) {
		result := corepack.Validate(root)
		return jsonrpc.Success(request.ID, PackValidateResult{Valid: result.Valid, Diagnostics: result.Diagnostics})
	}
	diagnostics := server.validatePackIfDeclared(root)
	result := corepack.Result{Valid: len(diagnostics) == 0, Diagnostics: diagnostics}
	return jsonrpc.Success(request.ID, PackValidateResult{Valid: result.Valid, Diagnostics: result.Diagnostics})
}
