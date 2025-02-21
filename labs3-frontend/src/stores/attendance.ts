import { defineStore } from 'pinia';
import { ILesson, IAttendance } from 'src/models/attendance/attendance';
import { IVisiting } from 'src/models/attendance/visiting';
import { AttendanceService } from 'src/services/attendance';
import { Ref, ref } from 'vue';

export const useAttendanceStore = defineStore('attendance', () => {
  const lessons: Ref<ILesson[] | null> = ref(null);
  const lessonsVisiting: Ref<IVisiting.StudentVisiting[] | null> = ref(null);
  const currentDiscipline: Ref<string> = ref('');
  const allDisciplineLessons = ref<ILesson[] | null>(null);
  const lessonsDate: Ref<IAttendance.LessonDate[] | null> = ref(null);
  const seminars: Ref<IAttendance.Seminar[] | null> = ref(null);
  const seminarVisiting: Ref<IVisiting.StudentVisiting[] | null> = ref(null);
  const groupLessons: Ref<IAttendance.LessonDate[] | null> = ref(null);

  const setCurrentDiscipline = (id: string) => {
    currentDiscipline.value = id;
  };

  const getLessons = async (disciplineId: string) => {
    const res = await AttendanceService.fetch(disciplineId);
    if (res.data) {
      lessons.value = res.data.lessons;
    }
  };

  const deleteLesson = async (id: string) => {
    await AttendanceService.deleteLesson(id);
  };

  const changeLessonName = async (lessonId: string, name: string) => {
    await AttendanceService.changeLessonName({
      lesson_id: lessonId,
      name,
    });
  };

  const createLesson = async (name: string) => {
    await AttendanceService.createLesson({
      discipline_id: Number(currentDiscipline.value),
      name,
    });
  };

  const getLessonsVisiting = async (groupId: number, lessonId: number) => {
    const res = await AttendanceService.getLessonVisiting({
      group_id: groupId,
      lesson_id: lessonId,
    });
    if (res.data) {
      lessonsVisiting.value = res.data.students;
    }
  };

  const changeLessonVisiting = async (
    isAbsent: boolean,
    lessonId: number,
    userId: number
  ) => {
    await AttendanceService.changeLessonVisiting({
      is_absent: isAbsent,
      lesson_id: lessonId,
      user_id: userId,
    });
  };

  const addLessonVisiting = async (
    isAbsent: boolean,
    lessonId: number,
    userId: number
  ) => {
    await AttendanceService.addLessonVisiting({
      is_absent: isAbsent,
      lesson_id: lessonId,
      user_id: userId,
    });
  };

  const getDisciplineLessons = async (disciplineId: number) => {
    const res = await AttendanceService.getDisciplineLessons(disciplineId);
    if (res.data) {
      allDisciplineLessons.value = res.data?.lessons;
    }
  };

  const getLessonDate = async (groupId: number, lessonId: number) => {
    await AttendanceService.getLessonDate({
      group_id: groupId,
      lesson_id: lessonId,
    });
  };

  const getSeminars = async (disciplineId: number, groupId: number) => {
    const res = await AttendanceService.fetchSeminars({
      discipline_id: disciplineId,
      group_id: groupId,
    });

    if (res.data) {
      seminars.value = res.data.seminars;
    }
  };

  const changeSeminarName = async (seminarId: number, name: string) => {
    await AttendanceService.changeSeminarName({
      seminar_id: seminarId,
      name,
    });
  };

  const changeSeminarDate = async (seminarId: number, date: number) => {
    await AttendanceService.changeSeminarDate({
      seminar_id: seminarId,
      date,
    });
  };

  const createSeminar = async (data: IAttendance.CreateSeminar) => {
    await AttendanceService.createSeminar(data);
  };

  const deleteSeminar = async (id: number) => {
    await AttendanceService.deleteSeminar(id);
  };

  const getGroupLessons = async (groupId: number, disciplineId: number) => {
    const res = await AttendanceService.getGroupLessons({
      group_id: groupId,
      discipline_id: disciplineId,
    });

    if (res.data) {
      groupLessons.value = res.data.lessons;
    }
  };

  const addLessonDate = async (data: IAttendance.AddLessonDate) => {
    await AttendanceService.addLessonDate(data);
  };

  const getSeminarVisiting = async (seminarId: number) => {
    const res = await AttendanceService.getSeminarVisiting(seminarId);
    if (res.data) {
      seminarVisiting.value = res.data.students;
    }
  };

  const addSeminarVisiting = async (
    isAbsent: boolean,
    seminarId: number,
    userId: number
  ) => {
    await AttendanceService.addSeminarVisiting({
      is_absent: isAbsent,
      seminar_id: seminarId,
      user_id: userId,
    });
  };

  const changeSeminarVisiting = async (
    isAbsent: boolean,
    userId: number,
    seminarId: number
  ) => {
    await AttendanceService.changeSeminarVisiting({
      is_absent: isAbsent,
      user_id: userId,
      seminar_id: seminarId,
    });
  };

  const changeLessonDate = async (
    date: number,
    groupId: number,
    lessonId: number
  ) => {
    await AttendanceService.changeLessonDate({
      lesson_id: lessonId,
      date,
      group_id: groupId,
    });
  };

  const deleteLessonDate = async (lessonId: number, groupId: number) => {
    await AttendanceService.deleteLessonDate({
      lesson_id: lessonId,
      group_id: groupId,
    });
  };

  return {
    lessons,
    lessonsVisiting,
    currentDiscipline,
    allDisciplineLessons,
    seminars,
    seminarVisiting,
    groupLessons,
    getLessons,
    deleteLesson,
    changeLessonName,
    createLesson,
    getLessonsVisiting,
    changeLessonVisiting,
    addLessonVisiting,
    setCurrentDiscipline,
    getDisciplineLessons,
    getLessonDate,
    getSeminars,
    createSeminar,
    changeSeminarName,
    changeSeminarDate,
    deleteSeminar,
    getGroupLessons,
    addLessonDate,
    getSeminarVisiting,
    addSeminarVisiting,
    changeSeminarVisiting,
    changeLessonDate,
    deleteLessonDate,
  };
});
