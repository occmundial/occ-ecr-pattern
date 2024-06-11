import { App, CfnOutput, Stack, assertions } from 'aws-cdk-lib'
import * as iam from 'aws-cdk-lib/aws-iam'
import { OCCEcrPattern } from '../src'

test('stack has required resources', () => {
  const app = new App()
  const stack = new Stack(app, 'ECRStack')

  const repository = new OCCEcrPattern(stack, 'ECRPattern', {
    imageName: 'test-repository',
    scanOnPush: true,
    principals: [new iam.ArnPrincipal('arn:aws:iam::12345678990:root')],
  })

  const template = assertions.Template.fromStack(stack)
  template.resourceCountIs('AWS::ECR::Repository', 1)

  new CfnOutput(stack, 'RepositoryName', {
    value: repository.ecrImageName,
  })

  new CfnOutput(stack, 'RepositoryArn', {
    value: repository.ecrArn,
  })
})
