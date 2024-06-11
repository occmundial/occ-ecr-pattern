import * as cdk from 'aws-cdk-lib'
import * as iam from 'aws-cdk-lib/aws-iam'

export interface IEcrProps extends cdk.StackProps {
  imageName: string
  scanOnPush: boolean
  principals: Array<iam.IPrincipal>
}
