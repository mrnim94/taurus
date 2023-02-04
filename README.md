# Taurus

## We are using Golang and AWS CLI to integrate with the AWS components.

### Auto Scaling Group.

You only need to create a configuration file in path config_file and the below content:

```yaml
autoscalings:
  - group_name: GroupName1
    profile: profile1
    schedule: "32 5 * * *"
    config:
      min: 1
      max: 1
      desired: 1
#  - group_name: GroupName2
#    profile: profile2
#    schedule: "*/2 * * * *"
#    config:
#      min: 1
#      max: 1
#      desired: 1
#  - group_name: GroupName3
#    profile: profile3
#    schedule: "*/3 * * * *"
#    config:
#      min: 1
#      max: 1
#      desired: 1
```

Taurus will update the auto-scaling group according to your wishes.