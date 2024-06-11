import * as cdk from 'aws-cdk-lib'
import * as ecr from 'aws-cdk-lib/aws-ecr'
import * as iam from 'aws-cdk-lib/aws-iam'
import { Construct } from 'constructs'

export interface EcrProps extends cdk.StackProps {
  imageName: string
  scanOnPush: boolean
  principals: Array<iam.IPrincipal>
}

export class OCCEcrPattern extends Construct {
  constructor(scope: Construct, id: string, props: EcrProps) {
    super(scope, id)

    const respository = new ecr.Repository(this, id, {
      imageScanOnPush: true,
      repositoryName: props.imageName,
    })

    respository.addToResourcePolicy(
      new iam.PolicyStatement({
        effect: iam.Effect.ALLOW,
        actions: [
          'ecr:GetDownloadUrlForLayer',
          'ecr:BatchGetImage',
          'ecr:BatchCheckLayerAvailability',
          'ecr:PutImage',
          'ecr:InitiateLayerUpload',
          'ecr:UploadLayerPart',
          'ecr:CompleteLayerUpload',
          'ecr:DescribeRepositories',
          'ecr:GetRepositoryPolicy',
          'ecr:ListImages',
          'ecr:DeleteRepository',
          'ecr:BatchDeleteImage',
          'ecr:SetRepositoryPolicy',
          'ecr:DeleteRepositoryPolicy',
        ],
        principals: props.principals,
      })
    )

    new cdk.CfnOutput(this, 'RepositoryArn', {
      value: respository.repositoryArn,
    })
  }
}
