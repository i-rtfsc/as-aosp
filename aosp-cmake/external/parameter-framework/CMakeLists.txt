# Copyright (c) 2024 anqi.huang@outlook.com
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use self file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

cmake_minimum_required(VERSION 3.5)

project(remote-process)

# Check if the major CMakeLists.txt file has correctly set
# the ANDROID_ROOT directory.
if (IS_DIRECTORY ${ANDROID_ROOT})
    message(STATUS "${PROJECT_NAME}, has been set android root")
else ()
    if (IS_DIRECTORY ${BUILD_NATIVE_ROOT})
        message(WARNING "${PROJECT_NAME}, set ANDROID_ROOT via BUILD_NATIVE_ROOT=${BUILD_NATIVE_ROOT}")
        set(ANDROID_ROOT ${BUILD_NATIVE_ROOT})
        set(OUT_INTERMEDIATES_ROOT ${ANDROID_ROOT}/out/soong/.intermediates)
    else ()
        message(FATAL_ERROR "${PROJECT_NAME}, ANDROID_ROOT was NULL")
    endif ()
endif ()

#################### remote-process ####################
set(FILES
        ${ANDROID_ROOT}/external/parameter-framework/upstream/remote-process/main.cpp
)
set(SOURCE_FILES "")
foreach (file ${FILES})
    if (EXISTS ${file})
        list(APPEND SOURCE_FILES ${file})
    endif ()
endforeach ()

if (NOT SOURCE_FILES)
    message(WARNING "${PROJECT_NAME}, source files not exist")
else ()
    add_executable(remote-process ${SOURCE_FILES})

    target_include_directories(remote-process
            PUBLIC
            ${ANDROID_ROOT}/external/parameter-framework/upstream/utility
            ${ANDROID_ROOT}/external/parameter-framework/support/android/asio
            ${ANDROID_ROOT}/external/parameter-framework/asio/include
            ${ANDROID_ROOT}/external/parameter-framework
            ${ANDROID_ROOT}/external/parameter-framework/upstream/remote-processor
            ${ANDROID_ROOT}/external/parameter-framework/support/android/remote-processor
    )

    # Add c/c++ flags
    target_compile_options(remote-process PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-local-typedef
    )

    target_link_libraries(remote-process
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libremote-processor
    )
endif ()
#################### remote-process ####################

#################### libpfw_utility ####################
set(FILES
        ${ANDROID_ROOT}/external/parameter-framework/upstream/utility/DynamicLibrary.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/utility/posix/DynamicLibrary.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/utility/Tokenizer.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/utility/Utility.cpp
)
set(SOURCE_FILES "")
foreach (file ${FILES})
    if (EXISTS ${file})
        list(APPEND SOURCE_FILES ${file})
    endif ()
endforeach ()

if (NOT SOURCE_FILES)
    message(WARNING "${PROJECT_NAME}, source files not exist")
else ()
    add_library(libpfw_utility SHARED ${SOURCE_FILES})

    target_include_directories(libpfw_utility
            PUBLIC
            ${ANDROID_ROOT}/external/parameter-framework/upstream/utility
            ${ANDROID_ROOT}/external/parameter-framework
    )

    # Add c/c++ flags
    target_compile_options(libpfw_utility PRIVATE
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(libpfw_utility
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libpfw_utility ####################

#################### libparameter ####################
set(FILES
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/ParameterMgrPlatformConnector.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/LoggingElementBuilderTemplate.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/StringParameterType.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/SyncerSet.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/BitParameter.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/BaseParameter.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/ParameterBlockType.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/FloatingPointParameterType.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/SelectionCriteriaDefinition.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/EnumValuePair.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/SelectionCriteria.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/SelectionCriterionRule.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/AreaConfiguration.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/BitParameterBlockType.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/ConfigurationAccessContext.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/BitwiseAreaConfiguration.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/ArrayParameter.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/ParameterBlackboard.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/InstanceConfigurableElement.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/LogarithmicParameterAdaptation.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/ConfigurableDomain.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/FormattedSubsystemObject.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/MappingData.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/SubsystemElementBuilder.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/BooleanParameterType.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/FixedPointParameterType.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/ComponentType.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/EnumParameterType.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/RuleParser.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/VirtualSubsystem.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/Element.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/ParameterFrameworkConfiguration.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/SelectionCriterionLibrary.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/StringParameter.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/CompoundRule.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/ConfigurableDomains.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/VirtualSyncer.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/MappingContext.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/LinearParameterAdaptation.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/ComponentLibrary.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/BitParameterBlock.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/ParameterMgrFullConnector.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/ConfigurableElement.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/ConfigurableElementAggregator.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/SubsystemObject.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/TypeElement.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/PathNavigator.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/ElementLocator.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/SimulatedBackSynchronizer.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/Parameter.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/ComponentInstance.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/InstanceDefinition.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/SubsystemObjectCreator.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/ParameterType.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/DomainConfiguration.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/PluginLocation.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/HardwareBackSynchronizer.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/SystemClass.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/ElementLibrary.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/ParameterAccessContext.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/XmlParameterSerializingContext.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/ElementHandle.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/ParameterMgr.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/SelectionCriterionType.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/Subsystem.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/BaseIntegerParameterType.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/BitParameterType.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/SelectionCriterion.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/XmlElementSerializingContext.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/ElementLibrarySet.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/FrameworkConfigurationLocation.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/ParameterAdaptation.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/XmlFileIncluderElement.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/xmlserializer/XmlElement.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/xmlserializer/XmlSerializingContext.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/xmlserializer/XmlMemoryDocSource.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/xmlserializer/XmlDocSource.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/xmlserializer/XmlMemoryDocSink.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/xmlserializer/XmlStreamDocSink.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/CommandHandlerWrapper.cpp
)
set(SOURCE_FILES "")
foreach (file ${FILES})
    if (EXISTS ${file})
        list(APPEND SOURCE_FILES ${file})
    endif ()
endforeach ()

if (NOT SOURCE_FILES)
    message(WARNING "${PROJECT_NAME}, source files not exist")
else ()
    add_library(libparameter SHARED ${SOURCE_FILES})

    target_include_directories(libparameter
            PUBLIC
            ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter
            ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/log/include
            ${ANDROID_ROOT}/external/parameter-framework/upstream/parameter/include
            ${ANDROID_ROOT}/external/parameter-framework/upstream/xmlserializer
            ${ANDROID_ROOT}/external/parameter-framework/upstream/remote-processor
            ${ANDROID_ROOT}/external/parameter-framework/support/android/parameter
            ${ANDROID_ROOT}/external/parameter-framework
            ${ANDROID_ROOT}/external/parameter-framework/upstream/utility
            ${ANDROID_ROOT}/external/parameter-framework/support/android/remote-processor
    )

    # Add c/c++ flags
    target_compile_options(libparameter PRIVATE
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(libparameter
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libpfw_utility
            libremote-processor
            libxml2
    )
endif ()
#################### libparameter ####################

#################### libremote-processor ####################
set(FILES
        ${ANDROID_ROOT}/external/parameter-framework/upstream/remote-processor/RequestMessage.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/remote-processor/Message.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/remote-processor/AnswerMessage.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/remote-processor/RemoteProcessorServer.cpp
        ${ANDROID_ROOT}/external/parameter-framework/upstream/remote-processor/BackgroundRemoteProcessorServer.cpp
)
set(SOURCE_FILES "")
foreach (file ${FILES})
    if (EXISTS ${file})
        list(APPEND SOURCE_FILES ${file})
    endif ()
endforeach ()

if (NOT SOURCE_FILES)
    message(WARNING "${PROJECT_NAME}, source files not exist")
else ()
    add_library(libremote-processor SHARED ${SOURCE_FILES})

    target_include_directories(libremote-processor
            PUBLIC
            ${ANDROID_ROOT}/external/parameter-framework/upstream/remote-processor
            ${ANDROID_ROOT}/external/parameter-framework/support/android/remote-processor
            ${ANDROID_ROOT}/external/parameter-framework/asio/include
            ${ANDROID_ROOT}/external/parameter-framework/support/android/asio
            ${ANDROID_ROOT}/external/parameter-framework
            ${ANDROID_ROOT}/external/parameter-framework/upstream/utility
    )

    # Add c/c++ flags
    target_compile_options(libremote-processor PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-local-typedef
    )

    target_link_libraries(libremote-processor
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libpfw_utility
    )
endif ()
#################### libremote-processor ####################
