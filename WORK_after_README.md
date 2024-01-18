项目文档地址：https://shimo.im/docs/8Nk6evmlaniX6wqL

go版本采用go1.19,开发之前可以使用go version命令查看自己的go版本

> 理论上go1.19的子版本都可以,稍高的版本应该也行,但是不建议用很新的版本

router.go中分出了2个api组，仅做示例，目的是区分各模块使用的api，使得结构更加清晰，更好维护，后续可以自己新增api组的个数，也可以修改api组的名字

public目录为服务器静态资源的存放位置

config目录为配置文件存放目录,其中config.go为全局配置文件

.gitignore文件也已经编写，clone后请将其中对于config文件的注释去掉，以免上传config文件(目前阶段没有服务器地址，似乎上传也行)

# Angular 规范

```txt
<type>(<scope>): <subject>
// 空一行
<body>
// 空一行
<footer>
```

样例

```txt
标题行：50个字符以内，描述主要变更内容

主体内容：更详细的说明文本，建议72个字符以内。 需要描述的信息包括: 
    - 为什么这个变更是必须的? 它可能是用来修复一个bug，增加一个feature，提升性能、可靠性、稳定性等等
    - 他如何解决这个问题? 具体描述解决问题的步骤
    - 是否存在副作用、风险? 

尾部：如果需要的化可以添加一个链接到issue地址或者其它文档，或者关闭某个issue
```

[Angular 规范详细文档](https://docs.google.com/document/d/1QrDFcIiPjSLDn3EL15IJygNPiHORgU1_OOAqWjiDU5Y)

## 参数说明

### Header

一行，三个字段

- type（必需）
- scope（可选）
- subject（必需）

#### type: 代表某次提交的类型，比如是修复一个bug还是增加一个新的feature

| type类型 | Annotation |
|--|--|
| feat	 | 新功能（feature） |
| fix	 | 修补bug |
| docs	 | 仅仅修改了文档，比如README, CHANGELOG, CONTRIBUTE等等 |
| style	 | 仅仅修改了空格、格式缩进、都好等等，不改变代码逻辑 |
| refactor | 代码重构，没有加新功能或者修复bug |
| test	 | 测试用例，包括单元测试、集成测试等 |
| chore	 | 改变构建流程、或者增加依赖库、工具等 |
| revert | 回滚到上一个版本 |

#### scope: 用于说明 commit 影响的范围，比如数据层、控制层、视图层等等，视项目不同而不同

#### subject: 是 commit 目的的简短描述，不超过50个字符

### Body部分

是对本次 commit 的详细描述，可以分成多行。

范例：

```txt
More detailed explanatory text, if necessary.  Wrap it to 
about 72 characters or so. 

Further paragraphs come after blank lines.

- Bullet points are okay, too
- Use a hanging indent
```

1. 使用第一人称现在时，比如使用change而不是changed或changes
2. 应该说明代码变动的动机，以及与以前行为的对比

### Footer部分

----