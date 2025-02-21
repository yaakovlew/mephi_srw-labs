import { defineStore } from 'pinia';
import { Ref, ref } from 'vue';
import { ITest } from 'src/models/test/test';
import { TestService } from 'src/services/test';
import { ITheme } from 'src/models/test/theme';
import { IQuestion } from 'src/models/test/question';
import { download } from 'src/utils/download';

export const useTestsStore = defineStore('tests', () => {
  const tests: Ref<ITest.Test[]> = ref([]);
  const openedTests: Ref<ITest.Test[]> = ref([]);
  const themes: Ref<ITheme.ThemeTest[]> = ref([]);
  const allQuestions: Ref<IQuestion.Question[]> = ref([]);
  const questions: Ref<IQuestion.Question[]> = ref([]);
  const allThemes: Ref<ITheme.Theme[]> = ref([]);
  const answers: Ref<IQuestion.Answer[]> = ref([]);
  const questionsWithoutTheme: Ref<IQuestion.Question[]> = ref([]);
  const sectionTests: Ref<ITest.Test[]> = ref([]);
  const test: Ref<IQuestion.StudentTest> = ref([]);
  const disciplineTests: Ref<ITest.Test[]> = ref([]);
  const openedTestsStudent: Ref<ITest.Test[]> = ref([]);
  const studentsOpenedTest: Ref<ITest.StudentOpenTest[]> = ref([]);
  const allSectionsTests: Ref<Record<number, ITest.Test[]>> = ref({});
  const studentsDoneTests: Ref<ITest.Test[]> = ref([]);

  const getTests = async () => {
    const res = await TestService.fetch();
    if (res.data) {
      tests.value = res.data.ru;
    }
  };

  const createTest = async (data: ITest.CreateTest) => {
    await TestService.createTest(data);
  };

  const closeTest = async (data: ITest.CloseTest) => {
    await TestService.closeTest(data);
  };

  const openTest = async (data: ITest.OpenTest) => {
    await TestService.openTest(data);
  };

  const getOpenedTest = async (data: ITest.GetOpenedTest) => {
    const res = await TestService.getOpenedTest(data);
    if (res.data) {
      openedTests.value = res.data;
    }
  };

  const getTestReport = async (
    data: ITest.GetTestReport,
    fileName = 'report.txt'
  ) => {
    const res = await TestService.getTestReport(data);
    download(fileName, res.data);
  };

  const changeTestDescription = async (data: ITest.ChangeTestDescription) => {
    await TestService.changeTestDescription(data);
  };

  const changeTestDuration = async (data: ITest.ChangeTestDuration) => {
    await TestService.changeTestDuration(data);
  };

  const changeTestName = async (data: ITest.ChangeTestName) => {
    await TestService.changeTestName(data);
  };

  const changeTestMark = async (data: ITest.ChangeTestMark) => {
    await TestService.changeTestMark(data);
  };

  const getTestMark = async (data: ITest.GetStudentMark) => {
    await TestService.getTestMark(data);
  };

  const getAllThemes = async () => {
    const res = await TestService.getAllThemes();
    if (res.data) {
      allThemes.value = res.data.themes;
    }
  };

  const createTheme = async (data: ITheme.CreateTheme) => {
    await TestService.createTheme(data);
  };

  const changeTestThemeName = async (data: ITheme.ChangeThemeName) => {
    await TestService.changeThemeName(data);
  };

  const getAllQuestions = async () => {
    const res = await TestService.getAllQuestions();
    if (res.data) {
      allQuestions.value = res.data.ru;
    }
  };

  const changeQuestion = async (data: IQuestion.ChangeQuestion) => {
    await TestService.changeQuestion(data);
  };

  const addQuestion = async (data: IQuestion.AddQuestionToTheme) => {
    await TestService.addQuestion(data);
  };

  const deleteQuestionFromTheme = async (
    data: IQuestion.RemoveQuestionFromTheme
  ) => {
    await TestService.deleteQuestionFromTheme(data);
  };

  const addAnswerToQuestion = async (data: IQuestion.AddAnswer) => {
    await TestService.addAnswerToQuestion(data);
  };

  const getAnswers = async (id: number) => {
    const res = await TestService.getAnswers(id);
    if (res.data) {
      answers.value = res.data.ru;
    }
  };

  const changeAnswerName = async (data: IQuestion.ChangeAnswerName) => {
    await TestService.changeAnswerName(data);
  };

  const changeAnswerIsRight = async (data: IQuestion.ChangeIsAnswerRight) => {
    await TestService.changeAnswerIsRight(data);
  };

  const deleteAnswer = async (id: number) => {
    await TestService.deleteAnswer(id);
  };

  const createQuestion = async (data: IQuestion.CreateQuestion) => {
    await TestService.createQuestion(data);
  };

  const getQuestions = async (id: number) => {
    const res = await TestService.getQuestions(id);
    if (res.data) {
      questions.value = res.data.ru;
    }
  };

  const deleteQuestion = async (id: number) => {
    await TestService.deleteQuestion(id);
  };

  const changeThemeWeight = async (data: ITheme.ChangeThemeWeight) => {
    await TestService.changeThemeWeight(data);
  };

  const getThemes = async (id: number) => {
    const res = await TestService.getThemes(id);
    if (res.data) {
      themes.value = res.data.themes;
    }
  };

  const deleteTheme = async (id: number) => {
    await TestService.deleteTheme(id);
  };

  const getQuestionsWithoutTheme = async () => {
    const res = await TestService.getQuestionsWithoutTheme();
    if (res.data) {
      questionsWithoutTheme.value = res.data.questions;
    }
  };

  const addThemeToTest = async (data: ITest.AddTheme) => {
    await TestService.addThemeToTest(data);
  };

  const deleteThemeFromTest = async (data: ITest.DeleteTheme) => {
    await TestService.deleteThemeFromTest(data);
  };

  const changeThemeCount = async (data: ITest.AddTheme) => {
    await TestService.changeThemeCount(data);
  };

  const getSectionTests = async (id: number) => {
    const res = await TestService.getSectionTests(id);
    if (res.data) {
      sectionTests.value = res.data.ru;
    }
    return { data: res.data, id };
  };

  const getAllSectionsTests = async (id: number[]) => {
    const promises: Promise<any>[] = [];
    id.forEach((id) => {
      promises.push(getSectionTests(id));
    });
    const res = await Promise.all(promises);
    allSectionsTests.value = {};
    if (res) {
      res.forEach((r) => {
        allSectionsTests.value[r.id] = r.data.ru;
      });
    }
  };

  const getTest = async (id: number) => {
    const res = await TestService.getTest(id);
    if (res.data) {
      test.value = res.data.ru;
    }
  };

  const getDisciplineTests = async (id: number) => {
    const res = await TestService.getDisciplineTests(id);
    if (res.data) {
      disciplineTests.value = res.data.ru;
    }
  };

  const getOpenedTests = async () => {
    const res = await TestService.getOpenedTests();
    if (res.data) {
      openedTestsStudent.value = res.data.ru;
    }
  };

  const testReport = async (id: number, fileName = 'report.txt') => {
    const res = await TestService.testReport(id);
    download(fileName, res.data);
  };

  const getStudentsOpenTest = async (id: number) => {
    const res = await TestService.getStudentsOpenTest(id);
    if (res.data) {
      studentsOpenedTest.value = res.data.students;
    }
  };

  const getDoneTestsStudent = async () => {
    const res = await TestService.getDoneTestsStudent();
    if (res.data) {
      studentsDoneTests.value = res.data.ru;
    }
  };

  return {
    getTests,
    tests,
    createTest,
    closeTest,
    openTest,
    getOpenedTest,
    openedTests,
    getTestReport,
    changeTestDescription,
    changeTestDuration,
    changeTestName,
    changeTestMark,
    getTestMark,
    getThemes,
    createTheme,
    changeTestThemeName,
    getAllQuestions,
    changeQuestion,
    addQuestion,
    deleteQuestionFromTheme,
    addAnswerToQuestion,
    getAnswers,
    changeAnswerName,
    changeAnswerIsRight,
    deleteAnswer,
    createQuestion,
    getQuestions,
    deleteQuestion,
    changeThemeWeight,
    getAllThemes,
    deleteTheme,
    themes,
    allQuestions,
    questions,
    allThemes,
    answers,
    getQuestionsWithoutTheme,
    questionsWithoutTheme,
    addThemeToTest,
    deleteThemeFromTest,
    changeThemeCount,
    getSectionTests,
    sectionTests,
    getTest,
    test,
    getDisciplineTests,
    disciplineTests,
    getOpenedTests,
    openedTestsStudent,
    testReport,
    getStudentsOpenTest,
    studentsOpenedTest,
    getDoneTestsStudent,
    getAllSectionsTests,
    allSectionsTests,
    studentsDoneTests,
  };
});
