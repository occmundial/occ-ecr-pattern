import * as ecr from 'aws-cdk-lib/aws-ecr'
import * as iam from 'aws-cdk-lib/aws-iam'
import { Construct } from 'constructs'

export interface IEcrProps {
  readonly imageName: string
  readonly scanOnPush: boolean
  readonly principals: Array<iam.IPrincipal>
}

export class OCCEcrPattern extends Construct {
  public ecrArn: string
  public ecrImageName: string
  public ecr: ecr.IRepository

  constructor(scope: Construct, id: string, props: IEcrProps) {
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

    this.ecrArn = respository.repositoryArn
    this.ecrImageName = respository.repositoryName
    this.ecr = respository
  }
}
