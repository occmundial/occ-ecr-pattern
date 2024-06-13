# API Reference <a name="API Reference" id="api-reference"></a>

## Constructs <a name="Constructs" id="Constructs"></a>

### OCCEcrPattern <a name="OCCEcrPattern" id="@occmundial/occ-ecr-pattern.OCCEcrPattern"></a>

#### Initializers <a name="Initializers" id="@occmundial/occ-ecr-pattern.OCCEcrPattern.Initializer"></a>

```typescript
import { OCCEcrPattern } from '@occmundial/occ-ecr-pattern'

new OCCEcrPattern(scope: Construct, id: string, props: IEcrProps)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@occmundial/occ-ecr-pattern.OCCEcrPattern.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | *No description.* |
| <code><a href="#@occmundial/occ-ecr-pattern.OCCEcrPattern.Initializer.parameter.id">id</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@occmundial/occ-ecr-pattern.OCCEcrPattern.Initializer.parameter.props">props</a></code> | <code><a href="#@occmundial/occ-ecr-pattern.IEcrProps">IEcrProps</a></code> | *No description.* |

---

##### `scope`<sup>Required</sup> <a name="scope" id="@occmundial/occ-ecr-pattern.OCCEcrPattern.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

---

##### `id`<sup>Required</sup> <a name="id" id="@occmundial/occ-ecr-pattern.OCCEcrPattern.Initializer.parameter.id"></a>

- *Type:* string

---

##### `props`<sup>Required</sup> <a name="props" id="@occmundial/occ-ecr-pattern.OCCEcrPattern.Initializer.parameter.props"></a>

- *Type:* <a href="#@occmundial/occ-ecr-pattern.IEcrProps">IEcrProps</a>

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@occmundial/occ-ecr-pattern.OCCEcrPattern.toString">toString</a></code> | Returns a string representation of this construct. |

---

##### `toString` <a name="toString" id="@occmundial/occ-ecr-pattern.OCCEcrPattern.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.

#### Static Functions <a name="Static Functions" id="Static Functions"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@occmundial/occ-ecr-pattern.OCCEcrPattern.isConstruct">isConstruct</a></code> | Checks if `x` is a construct. |

---

##### `isConstruct` <a name="isConstruct" id="@occmundial/occ-ecr-pattern.OCCEcrPattern.isConstruct"></a>

```typescript
import { OCCEcrPattern } from '@occmundial/occ-ecr-pattern'

OCCEcrPattern.isConstruct(x: any)
```

Checks if `x` is a construct.

Use this method instead of `instanceof` to properly detect `Construct`
instances, even when the construct library is symlinked.

Explanation: in JavaScript, multiple copies of the `constructs` library on
disk are seen as independent, completely different libraries. As a
consequence, the class `Construct` in each copy of the `constructs` library
is seen as a different class, and an instance of one class will not test as
`instanceof` the other class. `npm install` will not create installations
like this, but users may manually symlink construct libraries together or
use a monorepo tool: in those cases, multiple copies of the `constructs`
library can be accidentally installed, and `instanceof` will behave
unpredictably. It is safest to avoid using `instanceof`, and using
this type-testing method instead.

###### `x`<sup>Required</sup> <a name="x" id="@occmundial/occ-ecr-pattern.OCCEcrPattern.isConstruct.parameter.x"></a>

- *Type:* any

Any object.

---

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@occmundial/occ-ecr-pattern.OCCEcrPattern.property.node">node</a></code> | <code>constructs.Node</code> | The tree node. |
| <code><a href="#@occmundial/occ-ecr-pattern.OCCEcrPattern.property.ecr">ecr</a></code> | <code>aws-cdk-lib.aws_ecr.IRepository</code> | *No description.* |
| <code><a href="#@occmundial/occ-ecr-pattern.OCCEcrPattern.property.ecrArn">ecrArn</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@occmundial/occ-ecr-pattern.OCCEcrPattern.property.ecrImageName">ecrImageName</a></code> | <code>string</code> | *No description.* |

---

##### `node`<sup>Required</sup> <a name="node" id="@occmundial/occ-ecr-pattern.OCCEcrPattern.property.node"></a>

```typescript
public readonly node: Node;
```

- *Type:* constructs.Node

The tree node.

---

##### `ecr`<sup>Required</sup> <a name="ecr" id="@occmundial/occ-ecr-pattern.OCCEcrPattern.property.ecr"></a>

```typescript
public readonly ecr: IRepository;
```

- *Type:* aws-cdk-lib.aws_ecr.IRepository

---

##### `ecrArn`<sup>Required</sup> <a name="ecrArn" id="@occmundial/occ-ecr-pattern.OCCEcrPattern.property.ecrArn"></a>

```typescript
public readonly ecrArn: string;
```

- *Type:* string

---

##### `ecrImageName`<sup>Required</sup> <a name="ecrImageName" id="@occmundial/occ-ecr-pattern.OCCEcrPattern.property.ecrImageName"></a>

```typescript
public readonly ecrImageName: string;
```

- *Type:* string

---




## Protocols <a name="Protocols" id="Protocols"></a>

### IEcrProps <a name="IEcrProps" id="@occmundial/occ-ecr-pattern.IEcrProps"></a>

- *Implemented By:* <a href="#@occmundial/occ-ecr-pattern.IEcrProps">IEcrProps</a>


#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@occmundial/occ-ecr-pattern.IEcrProps.property.imageName">imageName</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@occmundial/occ-ecr-pattern.IEcrProps.property.principals">principals</a></code> | <code>aws-cdk-lib.aws_iam.IPrincipal[]</code> | *No description.* |
| <code><a href="#@occmundial/occ-ecr-pattern.IEcrProps.property.scanOnPush">scanOnPush</a></code> | <code>boolean</code> | *No description.* |

---

##### `imageName`<sup>Required</sup> <a name="imageName" id="@occmundial/occ-ecr-pattern.IEcrProps.property.imageName"></a>

```typescript
public readonly imageName: string;
```

- *Type:* string

---

##### `principals`<sup>Required</sup> <a name="principals" id="@occmundial/occ-ecr-pattern.IEcrProps.property.principals"></a>

```typescript
public readonly principals: IPrincipal[];
```

- *Type:* aws-cdk-lib.aws_iam.IPrincipal[]

---

##### `scanOnPush`<sup>Required</sup> <a name="scanOnPush" id="@occmundial/occ-ecr-pattern.IEcrProps.property.scanOnPush"></a>

```typescript
public readonly scanOnPush: boolean;
```

- *Type:* boolean

---

