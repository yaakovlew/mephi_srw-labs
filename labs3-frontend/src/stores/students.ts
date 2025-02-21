import { defineStore } from 'pinia';
import { IStudent } from 'src/models/student/student';
import { StudentService } from 'src/services/students';
import { Ref, ref } from 'vue';

export const useStudentsStore = defineStore('students', () => {
  const students: Ref<IStudent[] | null> = ref(null);

  const getStudents = async () => {
    const res = await StudentService.fetch();
    if (res.data) {
      students.value = res.data.students;
    }
  };

  const changeStudentData = async (studentId: string, groupId: string) => {
    await StudentService.changeGroup({
      user_id: studentId,
      group_id: groupId,
    });
  };

  const changeStudentPassword = async (
    studentId: string,
    newPassword: string
  ) => {
    await StudentService.changePassword({
      user_id: studentId,
      new_password: newPassword,
    });
  };

  return {
    students,
    getStudents,
    changeStudentData,
    changeStudentPassword,
  };
});
