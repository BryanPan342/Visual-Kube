// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package devicefarmiface provides an interface for the AWS Device Farm.
package devicefarmiface

import (
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/devicefarm"
)

// DeviceFarmAPI is the interface type for devicefarm.DeviceFarm.
type DeviceFarmAPI interface {
	CreateDevicePoolRequest(*devicefarm.CreateDevicePoolInput) (*request.Request, *devicefarm.CreateDevicePoolOutput)

	CreateDevicePool(*devicefarm.CreateDevicePoolInput) (*devicefarm.CreateDevicePoolOutput, error)

	CreateProjectRequest(*devicefarm.CreateProjectInput) (*request.Request, *devicefarm.CreateProjectOutput)

	CreateProject(*devicefarm.CreateProjectInput) (*devicefarm.CreateProjectOutput, error)

	CreateUploadRequest(*devicefarm.CreateUploadInput) (*request.Request, *devicefarm.CreateUploadOutput)

	CreateUpload(*devicefarm.CreateUploadInput) (*devicefarm.CreateUploadOutput, error)

	DeleteDevicePoolRequest(*devicefarm.DeleteDevicePoolInput) (*request.Request, *devicefarm.DeleteDevicePoolOutput)

	DeleteDevicePool(*devicefarm.DeleteDevicePoolInput) (*devicefarm.DeleteDevicePoolOutput, error)

	DeleteProjectRequest(*devicefarm.DeleteProjectInput) (*request.Request, *devicefarm.DeleteProjectOutput)

	DeleteProject(*devicefarm.DeleteProjectInput) (*devicefarm.DeleteProjectOutput, error)

	DeleteRunRequest(*devicefarm.DeleteRunInput) (*request.Request, *devicefarm.DeleteRunOutput)

	DeleteRun(*devicefarm.DeleteRunInput) (*devicefarm.DeleteRunOutput, error)

	DeleteUploadRequest(*devicefarm.DeleteUploadInput) (*request.Request, *devicefarm.DeleteUploadOutput)

	DeleteUpload(*devicefarm.DeleteUploadInput) (*devicefarm.DeleteUploadOutput, error)

	GetAccountSettingsRequest(*devicefarm.GetAccountSettingsInput) (*request.Request, *devicefarm.GetAccountSettingsOutput)

	GetAccountSettings(*devicefarm.GetAccountSettingsInput) (*devicefarm.GetAccountSettingsOutput, error)

	GetDeviceRequest(*devicefarm.GetDeviceInput) (*request.Request, *devicefarm.GetDeviceOutput)

	GetDevice(*devicefarm.GetDeviceInput) (*devicefarm.GetDeviceOutput, error)

	GetDevicePoolRequest(*devicefarm.GetDevicePoolInput) (*request.Request, *devicefarm.GetDevicePoolOutput)

	GetDevicePool(*devicefarm.GetDevicePoolInput) (*devicefarm.GetDevicePoolOutput, error)

	GetDevicePoolCompatibilityRequest(*devicefarm.GetDevicePoolCompatibilityInput) (*request.Request, *devicefarm.GetDevicePoolCompatibilityOutput)

	GetDevicePoolCompatibility(*devicefarm.GetDevicePoolCompatibilityInput) (*devicefarm.GetDevicePoolCompatibilityOutput, error)

	GetJobRequest(*devicefarm.GetJobInput) (*request.Request, *devicefarm.GetJobOutput)

	GetJob(*devicefarm.GetJobInput) (*devicefarm.GetJobOutput, error)

	GetProjectRequest(*devicefarm.GetProjectInput) (*request.Request, *devicefarm.GetProjectOutput)

	GetProject(*devicefarm.GetProjectInput) (*devicefarm.GetProjectOutput, error)

	GetRunRequest(*devicefarm.GetRunInput) (*request.Request, *devicefarm.GetRunOutput)

	GetRun(*devicefarm.GetRunInput) (*devicefarm.GetRunOutput, error)

	GetSuiteRequest(*devicefarm.GetSuiteInput) (*request.Request, *devicefarm.GetSuiteOutput)

	GetSuite(*devicefarm.GetSuiteInput) (*devicefarm.GetSuiteOutput, error)

	GetTestRequest(*devicefarm.GetTestInput) (*request.Request, *devicefarm.GetTestOutput)

	GetTest(*devicefarm.GetTestInput) (*devicefarm.GetTestOutput, error)

	GetUploadRequest(*devicefarm.GetUploadInput) (*request.Request, *devicefarm.GetUploadOutput)

	GetUpload(*devicefarm.GetUploadInput) (*devicefarm.GetUploadOutput, error)

	ListArtifactsRequest(*devicefarm.ListArtifactsInput) (*request.Request, *devicefarm.ListArtifactsOutput)

	ListArtifacts(*devicefarm.ListArtifactsInput) (*devicefarm.ListArtifactsOutput, error)

	ListArtifactsPages(*devicefarm.ListArtifactsInput, func(*devicefarm.ListArtifactsOutput, bool) bool) error

	ListDevicePoolsRequest(*devicefarm.ListDevicePoolsInput) (*request.Request, *devicefarm.ListDevicePoolsOutput)

	ListDevicePools(*devicefarm.ListDevicePoolsInput) (*devicefarm.ListDevicePoolsOutput, error)

	ListDevicePoolsPages(*devicefarm.ListDevicePoolsInput, func(*devicefarm.ListDevicePoolsOutput, bool) bool) error

	ListDevicesRequest(*devicefarm.ListDevicesInput) (*request.Request, *devicefarm.ListDevicesOutput)

	ListDevices(*devicefarm.ListDevicesInput) (*devicefarm.ListDevicesOutput, error)

	ListDevicesPages(*devicefarm.ListDevicesInput, func(*devicefarm.ListDevicesOutput, bool) bool) error

	ListJobsRequest(*devicefarm.ListJobsInput) (*request.Request, *devicefarm.ListJobsOutput)

	ListJobs(*devicefarm.ListJobsInput) (*devicefarm.ListJobsOutput, error)

	ListJobsPages(*devicefarm.ListJobsInput, func(*devicefarm.ListJobsOutput, bool) bool) error

	ListProjectsRequest(*devicefarm.ListProjectsInput) (*request.Request, *devicefarm.ListProjectsOutput)

	ListProjects(*devicefarm.ListProjectsInput) (*devicefarm.ListProjectsOutput, error)

	ListProjectsPages(*devicefarm.ListProjectsInput, func(*devicefarm.ListProjectsOutput, bool) bool) error

	ListRunsRequest(*devicefarm.ListRunsInput) (*request.Request, *devicefarm.ListRunsOutput)

	ListRuns(*devicefarm.ListRunsInput) (*devicefarm.ListRunsOutput, error)

	ListRunsPages(*devicefarm.ListRunsInput, func(*devicefarm.ListRunsOutput, bool) bool) error

	ListSamplesRequest(*devicefarm.ListSamplesInput) (*request.Request, *devicefarm.ListSamplesOutput)

	ListSamples(*devicefarm.ListSamplesInput) (*devicefarm.ListSamplesOutput, error)

	ListSamplesPages(*devicefarm.ListSamplesInput, func(*devicefarm.ListSamplesOutput, bool) bool) error

	ListSuitesRequest(*devicefarm.ListSuitesInput) (*request.Request, *devicefarm.ListSuitesOutput)

	ListSuites(*devicefarm.ListSuitesInput) (*devicefarm.ListSuitesOutput, error)

	ListSuitesPages(*devicefarm.ListSuitesInput, func(*devicefarm.ListSuitesOutput, bool) bool) error

	ListTestsRequest(*devicefarm.ListTestsInput) (*request.Request, *devicefarm.ListTestsOutput)

	ListTests(*devicefarm.ListTestsInput) (*devicefarm.ListTestsOutput, error)

	ListTestsPages(*devicefarm.ListTestsInput, func(*devicefarm.ListTestsOutput, bool) bool) error

	ListUniqueProblemsRequest(*devicefarm.ListUniqueProblemsInput) (*request.Request, *devicefarm.ListUniqueProblemsOutput)

	ListUniqueProblems(*devicefarm.ListUniqueProblemsInput) (*devicefarm.ListUniqueProblemsOutput, error)

	ListUniqueProblemsPages(*devicefarm.ListUniqueProblemsInput, func(*devicefarm.ListUniqueProblemsOutput, bool) bool) error

	ListUploadsRequest(*devicefarm.ListUploadsInput) (*request.Request, *devicefarm.ListUploadsOutput)

	ListUploads(*devicefarm.ListUploadsInput) (*devicefarm.ListUploadsOutput, error)

	ListUploadsPages(*devicefarm.ListUploadsInput, func(*devicefarm.ListUploadsOutput, bool) bool) error

	ScheduleRunRequest(*devicefarm.ScheduleRunInput) (*request.Request, *devicefarm.ScheduleRunOutput)

	ScheduleRun(*devicefarm.ScheduleRunInput) (*devicefarm.ScheduleRunOutput, error)

	UpdateDevicePoolRequest(*devicefarm.UpdateDevicePoolInput) (*request.Request, *devicefarm.UpdateDevicePoolOutput)

	UpdateDevicePool(*devicefarm.UpdateDevicePoolInput) (*devicefarm.UpdateDevicePoolOutput, error)

	UpdateProjectRequest(*devicefarm.UpdateProjectInput) (*request.Request, *devicefarm.UpdateProjectOutput)

	UpdateProject(*devicefarm.UpdateProjectInput) (*devicefarm.UpdateProjectOutput, error)
}

var _ DeviceFarmAPI = (*devicefarm.DeviceFarm)(nil)
