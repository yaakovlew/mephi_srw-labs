import { defineStore } from 'pinia';
import { Ref, ref } from 'vue';
import { IMark } from 'src/models/mark/mark';
import { MarksService } from 'src/services/mark';

export const useMarksStore = defineStore('marks', () => {
  const examMarks: Ref<IMark.Mark[]> = ref([]);
  const testMarks: Ref<IMark.Mark[]> = ref([]);
  const labaratoryMarks: Ref<IMark.Mark[]> = ref([]);
  const allTestsMarks: Ref<Record<number, IMark.Mark[]>> = ref({});
  const testsMarkStudent: Ref<IMark.TestMark[]> = ref([]);

  const getExamMarks = async (data: IMark.GetMarkExam) => {
    const res = await MarksService.getExamMark(data);
    if (res.data) {
      examMarks.value = res.data.marks;
    }
  };

  const getTestMarks = async (data: IMark.GetMarkTest) => {
    const res = await MarksService.getTestMark(data);
    if (res.data) {
      testMarks.value = res.data.marks;
    }
    return { data: res.data, id: data.test_id };
  };

  const getTestsMarks = async (id: number[], groupId: number) => {
    const promises: Promise<any>[] = [];
    id.forEach((id) => {
      promises.push(getTestMarks({ group_id: groupId, test_id: id }));
    });
    const res = await Promise.all(promises);
    allTestsMarks.value = {};
    res.forEach((r) => {
      allTestsMarks.value[r.id] = r.data.marks;
    });
  };

  const getLabaratoryMarks = async (data: IMark.GetMarkLabaratory) => {
    const res = await MarksService.getLabaratoryMark(data);
    if (res.data) {
      labaratoryMarks.value = res.data.marks;
    }
  };

  const postLabaratoryMark = async (data: IMark.PostMarkLabaratory) => {
    await MarksService.postLabaratoryMark(data);
  };

  const postTestMark = async (data: IMark.PostMarkTest) => {
    await MarksService.postTestMark(data);
  };

  const postExamMark = async (data: IMark.PostMarkExam) => {
    await MarksService.postExamMark(data);
  };

  const getTestMarksStudent = async (id: number) => {
    await MarksService.getTestsMarks(id);
  };

  const getDoneTestMarksStudent = async (ids: number[]) => {
    const promises: Promise<any>[] = [];
    ids.forEach((id) => {
      promises.push(MarksService.getTestMarkStudent(id));
    });
    const res = await Promise.all(promises);
    testsMarkStudent.value = res.map((test) => test.data);
  };

  return {
    examMarks,
    testMarks,
    labaratoryMarks,
    getExamMarks,
    getTestMarks,
    getLabaratoryMarks,
    postLabaratoryMark,
    postTestMark,
    postExamMark,
    getTestMarksStudent,
    getTestsMarks,
    allTestsMarks,
    getDoneTestMarksStudent,
    testsMarkStudent,
  };
});
