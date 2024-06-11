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
  name: 'occ-ecr-construct',
  projenrcTs: true,
  repositoryUrl: 'https://github.com/occmundial/occ-ecr-pattern.git',
  docgenFilePath: 'docs',
  readme: { filename: 'docs/README.md' },
  copyrightOwner: 'Online Career Center Mexico',
  copyrightPeriod: '2024',

  // deps: [],                /* Runtime dependencies of this module. */
  description:
    'OCC Pattern to create a ecr based on OCC best practices' /* The description is just a string that helps people understand the purpose of the package. */,
  // devDeps: [],             /* Build dependencies for this module. */
  packageName: 'occ-ecr-pattern' /* The "name" in package.json. */,
  publishToGo: {
    moduleName: 'github.com/occmundial/occ-ecr-pattern',
  },
  publishToPypi: {
    distName: 'occ.ecr_pattern',
    module: 'occ.ecr_pattern',
  },
  publishToNuget: {
    dotNetNamespace: 'Occ.EcrPattern',
    packageId: 'Occ.EcrPattern',
  },
})
project.synth()
