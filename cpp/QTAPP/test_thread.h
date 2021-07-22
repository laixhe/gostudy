#ifndef TEST_THREAD_H
#define TEST_THREAD_H

#include <QDialog>

namespace Ui {
    class TestThread;
}

class TestThread : public QDialog{
    Q_OBJECT

public:
    explicit TestThread(QWidget *parent = nullptr);
    ~TestThread();

signals:
    // 用于子线程 循环随机数的数量
    void starting(int num);

private:
    Ui::TestThread *ui;

};

#endif // TEST_THREAD_H
