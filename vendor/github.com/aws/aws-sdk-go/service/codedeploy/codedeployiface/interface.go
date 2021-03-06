// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package codedeployiface provides an interface for the AWS CodeDeploy.
package codedeployiface

import (
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/codedeploy"
)

// CodeDeployAPI is the interface type for codedeploy.CodeDeploy.
type CodeDeployAPI interface {
	AddTagsToOnPremisesInstancesRequest(*codedeploy.AddTagsToOnPremisesInstancesInput) (*request.Request, *codedeploy.AddTagsToOnPremisesInstancesOutput)

	AddTagsToOnPremisesInstances(*codedeploy.AddTagsToOnPremisesInstancesInput) (*codedeploy.AddTagsToOnPremisesInstancesOutput, error)

	BatchGetApplicationsRequest(*codedeploy.BatchGetApplicationsInput) (*request.Request, *codedeploy.BatchGetApplicationsOutput)

	BatchGetApplications(*codedeploy.BatchGetApplicationsInput) (*codedeploy.BatchGetApplicationsOutput, error)

	BatchGetDeploymentsRequest(*codedeploy.BatchGetDeploymentsInput) (*request.Request, *codedeploy.BatchGetDeploymentsOutput)

	BatchGetDeployments(*codedeploy.BatchGetDeploymentsInput) (*codedeploy.BatchGetDeploymentsOutput, error)

	BatchGetOnPremisesInstancesRequest(*codedeploy.BatchGetOnPremisesInstancesInput) (*request.Request, *codedeploy.BatchGetOnPremisesInstancesOutput)

	BatchGetOnPremisesInstances(*codedeploy.BatchGetOnPremisesInstancesInput) (*codedeploy.BatchGetOnPremisesInstancesOutput, error)

	CreateApplicationRequest(*codedeploy.CreateApplicationInput) (*request.Request, *codedeploy.CreateApplicationOutput)

	CreateApplication(*codedeploy.CreateApplicationInput) (*codedeploy.CreateApplicationOutput, error)

	CreateDeploymentRequest(*codedeploy.CreateDeploymentInput) (*request.Request, *codedeploy.CreateDeploymentOutput)

	CreateDeployment(*codedeploy.CreateDeploymentInput) (*codedeploy.CreateDeploymentOutput, error)

	CreateDeploymentConfigRequest(*codedeploy.CreateDeploymentConfigInput) (*request.Request, *codedeploy.CreateDeploymentConfigOutput)

	CreateDeploymentConfig(*codedeploy.CreateDeploymentConfigInput) (*codedeploy.CreateDeploymentConfigOutput, error)

	CreateDeploymentGroupRequest(*codedeploy.CreateDeploymentGroupInput) (*request.Request, *codedeploy.CreateDeploymentGroupOutput)

	CreateDeploymentGroup(*codedeploy.CreateDeploymentGroupInput) (*codedeploy.CreateDeploymentGroupOutput, error)

	DeleteApplicationRequest(*codedeploy.DeleteApplicationInput) (*request.Request, *codedeploy.DeleteApplicationOutput)

	DeleteApplication(*codedeploy.DeleteApplicationInput) (*codedeploy.DeleteApplicationOutput, error)

	DeleteDeploymentConfigRequest(*codedeploy.DeleteDeploymentConfigInput) (*request.Request, *codedeploy.DeleteDeploymentConfigOutput)

	DeleteDeploymentConfig(*codedeploy.DeleteDeploymentConfigInput) (*codedeploy.DeleteDeploymentConfigOutput, error)

	DeleteDeploymentGroupRequest(*codedeploy.DeleteDeploymentGroupInput) (*request.Request, *codedeploy.DeleteDeploymentGroupOutput)

	DeleteDeploymentGroup(*codedeploy.DeleteDeploymentGroupInput) (*codedeploy.DeleteDeploymentGroupOutput, error)

	DeregisterOnPremisesInstanceRequest(*codedeploy.DeregisterOnPremisesInstanceInput) (*request.Request, *codedeploy.DeregisterOnPremisesInstanceOutput)

	DeregisterOnPremisesInstance(*codedeploy.DeregisterOnPremisesInstanceInput) (*codedeploy.DeregisterOnPremisesInstanceOutput, error)

	GetApplicationRequest(*codedeploy.GetApplicationInput) (*request.Request, *codedeploy.GetApplicationOutput)

	GetApplication(*codedeploy.GetApplicationInput) (*codedeploy.GetApplicationOutput, error)

	GetApplicationRevisionRequest(*codedeploy.GetApplicationRevisionInput) (*request.Request, *codedeploy.GetApplicationRevisionOutput)

	GetApplicationRevision(*codedeploy.GetApplicationRevisionInput) (*codedeploy.GetApplicationRevisionOutput, error)

	GetDeploymentRequest(*codedeploy.GetDeploymentInput) (*request.Request, *codedeploy.GetDeploymentOutput)

	GetDeployment(*codedeploy.GetDeploymentInput) (*codedeploy.GetDeploymentOutput, error)

	GetDeploymentConfigRequest(*codedeploy.GetDeploymentConfigInput) (*request.Request, *codedeploy.GetDeploymentConfigOutput)

	GetDeploymentConfig(*codedeploy.GetDeploymentConfigInput) (*codedeploy.GetDeploymentConfigOutput, error)

	GetDeploymentGroupRequest(*codedeploy.GetDeploymentGroupInput) (*request.Request, *codedeploy.GetDeploymentGroupOutput)

	GetDeploymentGroup(*codedeploy.GetDeploymentGroupInput) (*codedeploy.GetDeploymentGroupOutput, error)

	GetDeploymentInstanceRequest(*codedeploy.GetDeploymentInstanceInput) (*request.Request, *codedeploy.GetDeploymentInstanceOutput)

	GetDeploymentInstance(*codedeploy.GetDeploymentInstanceInput) (*codedeploy.GetDeploymentInstanceOutput, error)

	GetOnPremisesInstanceRequest(*codedeploy.GetOnPremisesInstanceInput) (*request.Request, *codedeploy.GetOnPremisesInstanceOutput)

	GetOnPremisesInstance(*codedeploy.GetOnPremisesInstanceInput) (*codedeploy.GetOnPremisesInstanceOutput, error)

	ListApplicationRevisionsRequest(*codedeploy.ListApplicationRevisionsInput) (*request.Request, *codedeploy.ListApplicationRevisionsOutput)

	ListApplicationRevisions(*codedeploy.ListApplicationRevisionsInput) (*codedeploy.ListApplicationRevisionsOutput, error)

	ListApplicationRevisionsPages(*codedeploy.ListApplicationRevisionsInput, func(*codedeploy.ListApplicationRevisionsOutput, bool) bool) error

	ListApplicationsRequest(*codedeploy.ListApplicationsInput) (*request.Request, *codedeploy.ListApplicationsOutput)

	ListApplications(*codedeploy.ListApplicationsInput) (*codedeploy.ListApplicationsOutput, error)

	ListApplicationsPages(*codedeploy.ListApplicationsInput, func(*codedeploy.ListApplicationsOutput, bool) bool) error

	ListDeploymentConfigsRequest(*codedeploy.ListDeploymentConfigsInput) (*request.Request, *codedeploy.ListDeploymentConfigsOutput)

	ListDeploymentConfigs(*codedeploy.ListDeploymentConfigsInput) (*codedeploy.ListDeploymentConfigsOutput, error)

	ListDeploymentConfigsPages(*codedeploy.ListDeploymentConfigsInput, func(*codedeploy.ListDeploymentConfigsOutput, bool) bool) error

	ListDeploymentGroupsRequest(*codedeploy.ListDeploymentGroupsInput) (*request.Request, *codedeploy.ListDeploymentGroupsOutput)

	ListDeploymentGroups(*codedeploy.ListDeploymentGroupsInput) (*codedeploy.ListDeploymentGroupsOutput, error)

	ListDeploymentGroupsPages(*codedeploy.ListDeploymentGroupsInput, func(*codedeploy.ListDeploymentGroupsOutput, bool) bool) error

	ListDeploymentInstancesRequest(*codedeploy.ListDeploymentInstancesInput) (*request.Request, *codedeploy.ListDeploymentInstancesOutput)

	ListDeploymentInstances(*codedeploy.ListDeploymentInstancesInput) (*codedeploy.ListDeploymentInstancesOutput, error)

	ListDeploymentInstancesPages(*codedeploy.ListDeploymentInstancesInput, func(*codedeploy.ListDeploymentInstancesOutput, bool) bool) error

	ListDeploymentsRequest(*codedeploy.ListDeploymentsInput) (*request.Request, *codedeploy.ListDeploymentsOutput)

	ListDeployments(*codedeploy.ListDeploymentsInput) (*codedeploy.ListDeploymentsOutput, error)

	ListDeploymentsPages(*codedeploy.ListDeploymentsInput, func(*codedeploy.ListDeploymentsOutput, bool) bool) error

	ListOnPremisesInstancesRequest(*codedeploy.ListOnPremisesInstancesInput) (*request.Request, *codedeploy.ListOnPremisesInstancesOutput)

	ListOnPremisesInstances(*codedeploy.ListOnPremisesInstancesInput) (*codedeploy.ListOnPremisesInstancesOutput, error)

	RegisterApplicationRevisionRequest(*codedeploy.RegisterApplicationRevisionInput) (*request.Request, *codedeploy.RegisterApplicationRevisionOutput)

	RegisterApplicationRevision(*codedeploy.RegisterApplicationRevisionInput) (*codedeploy.RegisterApplicationRevisionOutput, error)

	RegisterOnPremisesInstanceRequest(*codedeploy.RegisterOnPremisesInstanceInput) (*request.Request, *codedeploy.RegisterOnPremisesInstanceOutput)

	RegisterOnPremisesInstance(*codedeploy.RegisterOnPremisesInstanceInput) (*codedeploy.RegisterOnPremisesInstanceOutput, error)

	RemoveTagsFromOnPremisesInstancesRequest(*codedeploy.RemoveTagsFromOnPremisesInstancesInput) (*request.Request, *codedeploy.RemoveTagsFromOnPremisesInstancesOutput)

	RemoveTagsFromOnPremisesInstances(*codedeploy.RemoveTagsFromOnPremisesInstancesInput) (*codedeploy.RemoveTagsFromOnPremisesInstancesOutput, error)

	StopDeploymentRequest(*codedeploy.StopDeploymentInput) (*request.Request, *codedeploy.StopDeploymentOutput)

	StopDeployment(*codedeploy.StopDeploymentInput) (*codedeploy.StopDeploymentOutput, error)

	UpdateApplicationRequest(*codedeploy.UpdateApplicationInput) (*request.Request, *codedeploy.UpdateApplicationOutput)

	UpdateApplication(*codedeploy.UpdateApplicationInput) (*codedeploy.UpdateApplicationOutput, error)

	UpdateDeploymentGroupRequest(*codedeploy.UpdateDeploymentGroupInput) (*request.Request, *codedeploy.UpdateDeploymentGroupOutput)

	UpdateDeploymentGroup(*codedeploy.UpdateDeploymentGroupInput) (*codedeploy.UpdateDeploymentGroupOutput, error)
}

var _ CodeDeployAPI = (*codedeploy.CodeDeploy)(nil)
