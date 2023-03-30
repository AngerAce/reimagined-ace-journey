# 定义学生类
class Student:
    def __init__(self, name, age, gender, grade, score):
        self.name = name
        self.age = age
        self.gender = gender
        self.grade = grade
        self.score = score

    def get_name(self):
        return self.name

    def get_age(self):
        return self.age

    def get_gender(self):
        return self.gender

    def get_grade(self):
        return self.grade

    def get_score(self):
        return self.score

    def set_name(self, name):
        self.name = name

    def set_age(self, age):
        self.age = age

    def set_gender(self, gender):
        self.gender = gender

    def set_grade(self, grade):
        self.grade = grade

    def set_score(self, score):
        self.score = score

# 定义学生信息管理类
class StudentManager:
    def __init__(self):
        self.students = []

    # 添加学生信息
    def add_student(self, student):
        self.students.append(student)

    # 删除学生信息
    def remove_student(self, index):
        del self.students[index]

    # 查看学生信息
    def view_students(self):
        print("Name\tAge\tGender\tGrade\tScore")
        for student in self.students:
            print(f"{student.get_name()}\t{student.get_age()}\t{student.get_gender()}\t{student.get_grade()}\t{student.get_score()}")

# 主程序
def main():
    # 创建学生信息管理对象
    manager = StudentManager()

    # 创建一个输入对象
    choice = 0
    while choice != 4:
        # 显示菜单
        print("1. Add a student")
        print("2. Remove a student")
        print("3. View students")
        print("4. Exit")
        choice = int(input("Enter your choice: "))

        if choice == 1:
            # 添加学生信息
            name = input("Enter student name: ")
            age = int(input("Enter student age: "))
            gender = input("Enter student gender: ")
            grade = int(input("Enter student grade: "))
            score = float(input("Enter student score: "))
            student = Student(name, age, gender, grade, score)
            manager.add_student(student)
            print("Student added!")

        elif choice == 2:
            # 删除学生信息
            index = int(input("Enter student index: "))
            if index >= 0 and index < len(manager.students):
                manager.remove_student(index)
                print("Student removed!")
            else:
                print("Invalid student index!")

        elif choice == 3:
            # 查看学生信息
            manager.view_students()

        elif choice == 4:
            # 退出程序
            print("Goodbye!")

        else:
            # 输入错误，重新显示菜单
            print("Invalid choice, please try again.")

# 调用主程序
if __name__ == '__main__':
    main()
