//
// 字符串通常指的的是：String 和 &str ( &'static str )
//
// String 类型实际上是 std::string::String 类型，且保证内部只保存标准的 UTF-8 文本，
//        在操作时又是以 char 类型进行操作的，char 类型大小为 4 个字节(32位)的 Unicode 码，
//        为了解决使用 Unicode 带来的麻烦，提供了将 utf-8 字节序列转化为类型 char 的 vector 的方法
//
// String 是一个结构体，由栈(24bytes)的空间保存指向堆空间的地址、容量、长度 ，
//        且在离开作用域的时候会调用 drop 方法来释放内存空间(类似析构函数)
//
// &str 是字符串片段(在运行时已知长度的文本)是无法修改的，一种比较特殊的数组片段(slice)，
//      由两个部分组成 内存地址和比特长度(多少bytes)

fn main() {
    // 通过字符串字面量构造 &str 类型的对象(不可变的)
    let s = "字符串片段";
    println!("s={}", s);
    
    // 创建字符串
    // let mut s1 = "字符串...".to_string();
    let mut s1 = String::from("字符串..."); // 在堆上分配内存
    // 追加字符串
    s1.push_str("追加字符串...");
    // 拷贝(克隆)字符串（深拷贝）
    let s2 = s1.clone();
    println!("s1={}", s1);
    println!("s2={}", s2);
    // 浅拷贝（拷贝了指向堆空间的地址、长度、容量），并所有权转移到 s3，而 s2 的生命周期结束
    let s3 = s2;   // move
    //println!("{}", s2); // error
    println!("s3={}", s3);
    
    // 将字符串转换为类型为 char 的 vector
    let chars = s3.chars().collect::<Vec<char>>();
    println!("chars.len={}", chars.len()); // 结果：14
    // 由于 char 为 4 字节长，我们可以将其转化为 u32
    println!("{}", chars[0] as u32);
    
    // 连接字符串
    let s4 = ["你好", "，", "世界", "！"].concat();
    let s5 = ["a", "b", "c"].join(",");
    println!("s4={}", s4); // 结果：你好，世界！
    println!("s5={}",s5); // 结果：a,b,c
}
