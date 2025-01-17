# plusitall 程序核心代码

### 单元测试

通常,单元测试文件会放在与被测试的源代码文件相同的目录下,只是文件名以 \_test.go 结尾。

这样做有几个好处:

文件位置一致性: 测试文件与被测试的源代码文件放在同一个目录下,可以更好地维护项目结构,提高代码的可读性和可维护性。
命名约定一致性: 使用 \_test.go 作为文件名后缀是 Go 语言的惯例做法,可以让代码更容易被其他 Go 开发者理解和识别。
构建和测试便利性: 当您运行 go test 命令时,Go 编译器会自动发现同目录下以 \_test.go 结尾的文件,并将它们作为测试文件进行编译和执行。这样可以提高开发效率。
隔离性: 将测试文件与源代码文件放在同一目录下,可以确保测试代码只访问同一包中的公开 API,而不会意外地访问到包外的实现细节。这有助于维护良好的代码抽象和封装。
