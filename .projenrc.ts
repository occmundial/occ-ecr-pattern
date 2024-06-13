import { awscdk } from 'projen'
const project = new awscdk.AwsCdkConstructLibrary({
  author: 'Alex Parra',
  authorAddress: 'jparra@occ.com.mx',
  cdkVersion: '2.143.0',
  constructsVersion: '10.3.0',
  defaultReleaseBranch: 'main',
  devContainer: true,
  eslint: false,
  github: true,
  jsiiVersion: '~5.4.0',
  name: 'occ-ecr-pattern',
  projenrcTs: true,
  repositoryUrl: 'https://github.com/occmundial/occ-ecr-pattern.git',
  readme: { filename: 'docs/README.md' },
  copyrightOwner: 'Online Career Center Mexico',
  copyrightPeriod: '2024',
  licensed: true,
  license: 'MIT',
  npmProvenance: false,

  // deps: [],
  description: 'OCC Pattern to create a ecr based on OCC best practices',
  // devDeps: [],
  packageName: 'occ-ecr-pattern',
  publishToGo: {
    moduleName: 'github.com/occmundial/occ-ecr-pattern',
  },
  publishToPypi: {
    distName: 'occ.ecr_pattern',
    module: 'occ.ecr_pattern',
  },
  publishToNuget: {
    dotNetNamespace: 'Occ.EcrPattern',
    packageId: 'OccEcrPattern',
  },
})
project.tasks
  .tryFind('release')
  ?.updateStep(4, { exec: 'git diff --ignore-space-at-eol --exit-code | tee' })

project.synth()
