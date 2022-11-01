// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: managementpb/backup/backups.proto

package backupv1

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	_ "google.golang.org/protobuf/types/known/wrapperspb"

	_ "github.com/percona/pmm/api/inventorypb"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = proto.Marshal
	_ = fmt.Errorf
	_ = math.Inf
)

func (this *StartBackupRequest) Validate() error {
	if this.ServiceId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ServiceId", fmt.Errorf(`value '%v' must not be an empty string`, this.ServiceId))
	}
	if this.LocationId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("LocationId", fmt.Errorf(`value '%v' must not be an empty string`, this.LocationId))
	}
	if this.RetryInterval != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.RetryInterval); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("RetryInterval", err)
		}
	}
	return nil
}

func (this *StartBackupResponse) Validate() error {
	return nil
}

func (this *ListArtifactCompatibleServicesRequest) Validate() error {
	if this.ArtifactId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ArtifactId", fmt.Errorf(`value '%v' must not be an empty string`, this.ArtifactId))
	}
	return nil
}

func (this *ListArtifactCompatibleServicesResponse) Validate() error {
	for _, item := range this.Mysql {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Mysql", err)
			}
		}
	}
	for _, item := range this.Mongodb {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Mongodb", err)
			}
		}
	}
	return nil
}

func (this *RestoreBackupRequest) Validate() error {
	if this.ServiceId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ServiceId", fmt.Errorf(`value '%v' must not be an empty string`, this.ServiceId))
	}
	if this.ArtifactId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ArtifactId", fmt.Errorf(`value '%v' must not be an empty string`, this.ArtifactId))
	}
	if this.PitrTimestamp != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.PitrTimestamp); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("PitrTimestamp", err)
		}
	}
	return nil
}

func (this *RestoreBackupResponse) Validate() error {
	return nil
}

func (this *ScheduledBackup) Validate() error {
	if this.StartTime != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.StartTime); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("StartTime", err)
		}
	}
	if this.RetryInterval != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.RetryInterval); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("RetryInterval", err)
		}
	}
	if this.LastRun != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.LastRun); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("LastRun", err)
		}
	}
	if this.NextRun != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.NextRun); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("NextRun", err)
		}
	}
	return nil
}

func (this *ScheduleBackupRequest) Validate() error {
	if this.ServiceId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ServiceId", fmt.Errorf(`value '%v' must not be an empty string`, this.ServiceId))
	}
	if this.LocationId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("LocationId", fmt.Errorf(`value '%v' must not be an empty string`, this.LocationId))
	}
	if this.CronExpression == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("CronExpression", fmt.Errorf(`value '%v' must not be an empty string`, this.CronExpression))
	}
	if this.StartTime != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.StartTime); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("StartTime", err)
		}
	}
	if this.RetryInterval != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.RetryInterval); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("RetryInterval", err)
		}
	}
	return nil
}

func (this *ScheduleBackupResponse) Validate() error {
	return nil
}

func (this *ListScheduledBackupsRequest) Validate() error {
	return nil
}

func (this *ListScheduledBackupsResponse) Validate() error {
	for _, item := range this.ScheduledBackups {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("ScheduledBackups", err)
			}
		}
	}
	return nil
}

func (this *ChangeScheduledBackupRequest) Validate() error {
	if this.ScheduledBackupId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ScheduledBackupId", fmt.Errorf(`value '%v' must not be an empty string`, this.ScheduledBackupId))
	}
	if this.Enabled != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Enabled); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Enabled", err)
		}
	}
	if this.CronExpression != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.CronExpression); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("CronExpression", err)
		}
	}
	if this.StartTime != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.StartTime); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("StartTime", err)
		}
	}
	if this.Name != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Name); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Name", err)
		}
	}
	if this.Description != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Description); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Description", err)
		}
	}
	if this.RetryInterval != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.RetryInterval); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("RetryInterval", err)
		}
	}
	if this.Retries != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Retries); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Retries", err)
		}
	}
	if this.Retention != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Retention); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Retention", err)
		}
	}
	return nil
}

func (this *ChangeScheduledBackupResponse) Validate() error {
	return nil
}

func (this *RemoveScheduledBackupRequest) Validate() error {
	if this.ScheduledBackupId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ScheduledBackupId", fmt.Errorf(`value '%v' must not be an empty string`, this.ScheduledBackupId))
	}
	return nil
}

func (this *RemoveScheduledBackupResponse) Validate() error {
	return nil
}

func (this *GetLogsRequest) Validate() error {
	if this.ArtifactId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ArtifactId", fmt.Errorf(`value '%v' must not be an empty string`, this.ArtifactId))
	}
	return nil
}

func (this *GetLogsResponse) Validate() error {
	for _, item := range this.Logs {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Logs", err)
			}
		}
	}
	return nil
}

func (this *LogChunk) Validate() error {
	return nil
}
