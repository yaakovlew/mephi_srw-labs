import { defineStore } from 'pinia';
import { Ref, ref } from 'vue';
import { IQuestion } from 'src/models/test/question';
import { TestService } from 'src/services/test';
import { ITest } from 'src/models/test/test';

export const useCurrentTestStore = defineStore('current-test', () => {
  const currentTest: Ref<number | null> = ref(
    localStorage.getItem('current-test') !== 'null'
      ? Number(localStorage.getItem('current-test'))
      : null
  );
  const test: Ref<IQuestion.StudentTest | null> = ref(null);

  const setCurrentTest = async (newTest: number | null) => {
    currentTest.value = newTest;
    localStorage.setItem('current-test', String(newTest));
  };

  const getTest = async () => {
    if (currentTest.value) {
      const res = await TestService.getTest(currentTest.value);
      if (res.data) {
        test.value = res.data.ru;
      }
    }
  };

  const passTest = async (data: ITest.PassTest) => {
    await TestService.passTest(data);
  };

  return { test, currentTest, setCurrentTest, getTest, passTest };
});
