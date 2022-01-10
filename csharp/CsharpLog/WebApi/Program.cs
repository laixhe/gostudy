using WebApi.Log;

var builder = WebApplication.CreateBuilder(args);

// ������־
builder.ConfigureLog();

// ע�� MVC �����������ֿ�����
builder.Services.AddControllers();

builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();

var app = builder.Build();

// �жϻ�������
if (app.Environment.IsDevelopment())
{
    app.UseSwagger();
    app.UseSwaggerUI();

    // �쳣�����м��
    app.UseDeveloperExceptionPage();
}

// �ɷ��ʺ�ʹ�þ�̬�ļ�
//app.UseStaticFiles();
// ��֤�м��
//app.UseAuthentication();
// ��Ȩ�м��
//app.UseAuthorization();
// �����м��
//app.UseCors();
// ȫ���쳣�����м��
//app.UseExceptionHandler();
// ����ͷ��Ϣת���м��
//app.UseForwardedHeaders();
// Session�м��
//app.UseSession();
// ע�������·�ɺ� MVC �м��
app.MapControllers();

// HTTP�������Ӧ��־�м��
app.UseHttpLogging();

app.Run();

// Mini API
// var builder = WebApplication.CreateBuilder(args);
// var app = builder.Build();
// app.MapGet("/", () => "Hello GET");
// app.MapPost("/", () => "Hello POST");
// app.MapPut("/", () => "Hello PUT");
// app.MapDelete("/", () => "Hello DELETE");
// app.MapGet("/user/{name}", (string name) => $"Hello {name}");
// app.Run();
