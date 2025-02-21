import { $apiLecturer, $apiStudent } from 'src/boot/axios';
import { ITest } from 'src/models/test/test';
import { ITheme } from 'src/models/test/theme';
import { useServiceAction } from 'src/utils/service/action';
import { IQuestion } from '../models/test/question';

export const TestService = {
  fetch: useServiceAction(() => $apiLecturer.get<ITest.GetTest>('/test')),
  createTest: useServiceAction((data: ITest.CreateTest) =>
    $apiLecturer.post('/test', data)
  ),
  closeTest: useServiceAction((data: ITest.CloseTest) =>
    $apiLecturer.delete(
      `/test/activity?test_id=${data.test_id}&user_id=${data.user_id}`
    )
  ),
  getOpenedTest: useServiceAction((data: ITest.GetOpenedTest) =>
    $apiLecturer.get<ITest.Test[]>(
      `/test&test_id=${data.test_id}?user_id=${data.user_id}`
    )
  ),
  openTest: useServiceAction((data: ITest.OpenTest) =>
    $apiLecturer.post('/test/activity', data)
  ),
  getTestReport: useServiceAction((data: ITest.GetTestReport) =>
    $apiLecturer.get(
      `/test/activity/report?test_id=${data.test_id}&user_id=${data.user_id}`
    )
  ),
  changeTestDescription: useServiceAction((data: ITest.ChangeTestDescription) =>
    $apiLecturer.put('/test/description', data)
  ),
  changeTestDuration: useServiceAction((data: ITest.ChangeTestDuration) =>
    $apiLecturer.put('/test/duration', data)
  ),
  changeTestName: useServiceAction((data: ITest.ChangeTestName) =>
    $apiLecturer.put('/test/name', data)
  ),
  changeTestMark: useServiceAction((data: ITest.ChangeTestMark) =>
    $apiLecturer.put('/test/mark', data)
  ),
  getTestMark: useServiceAction((data: ITest.GetStudentMark) =>
    $apiLecturer.get(`/test/mark/${data.test_id}/${data.user_id}`)
  ),
  getAllThemes: useServiceAction(() =>
    $apiLecturer.get<ITheme.GetThemes>('/test/theme')
  ),
  createTheme: useServiceAction((data: ITheme.CreateTheme) =>
    $apiLecturer.post('/test/theme', data)
  ),
  changeThemeName: useServiceAction((data: ITheme.ChangeThemeName) =>
    $apiLecturer.put('/test/theme/name', data)
  ),
  getAllQuestions: useServiceAction(() =>
    $apiLecturer.get<IQuestion.GetQuestions>('/test/theme/question')
  ),
  changeQuestion: useServiceAction((data: IQuestion.ChangeQuestion) =>
    $apiLecturer.put('/test/theme/question', data)
  ),
  addQuestion: useServiceAction((data: IQuestion.AddQuestionToTheme) =>
    $apiLecturer.post('/test/theme/question', data)
  ),
  deleteQuestionFromTheme: useServiceAction(
    (data: IQuestion.RemoveQuestionFromTheme) =>
      $apiLecturer.delete(
        `/test/theme/question?question_id=${data.question_id}&theme_id=${data.theme_id}`
      )
  ),
  addAnswerToQuestion: useServiceAction((data: IQuestion.AddAnswer) =>
    $apiLecturer.post('/test/theme/question/answer', data)
  ),
  changeAnswerName: useServiceAction((data: IQuestion.ChangeAnswerName) =>
    $apiLecturer.put('/test/theme/question/answer/name', data)
  ),
  changeAnswerIsRight: useServiceAction((data: IQuestion.ChangeIsAnswerRight) =>
    $apiLecturer.put('/test/theme/question/answer/right', data)
  ),
  getAnswers: useServiceAction((id: number) =>
    $apiLecturer.get<IQuestion.GetAnswers>(`/test/theme/question/answer/${id}`)
  ),
  deleteAnswer: useServiceAction((id: number) =>
    $apiLecturer.delete(`/test/theme/question/answer/${id}`)
  ),
  createQuestion: useServiceAction((data: IQuestion.CreateQuestion) =>
    $apiLecturer.post('/test/theme/question/create', data)
  ),
  getQuestions: useServiceAction((id: number) =>
    $apiLecturer.get<IQuestion.GetQuestions>(`/test/theme/question/${id}`)
  ),
  deleteQuestion: useServiceAction((id: number) =>
    $apiLecturer.delete(`/test/theme/question/${id}`)
  ),
  changeThemeWeight: useServiceAction((data: ITheme.ChangeThemeWeight) =>
    $apiLecturer.put('/test/theme/weight', data)
  ),
  getThemes: useServiceAction((id: number) =>
    $apiLecturer.get<ITheme.GetThemesTest>(`/test/theme/${id}`)
  ),
  deleteTheme: useServiceAction((id: number) =>
    $apiLecturer.delete(`/test/theme/${id}`)
  ),
  getQuestionsWithoutTheme: useServiceAction(() =>
    $apiLecturer.get<IQuestion.GetQuestionsWithotTheme>(
      '/test/theme/question/without-theme'
    )
  ),
  addThemeToTest: useServiceAction((data: ITest.AddTheme) =>
    $apiLecturer.post('/test/theme/add', data)
  ),
  deleteThemeFromTest: useServiceAction((data: ITest.DeleteTheme) =>
    $apiLecturer.delete(
      `/test/theme?test_id=${data.test_id}&theme_id=${data.theme_id}`
    )
  ),
  changeThemeCount: useServiceAction((data: ITest.AddTheme) =>
    $apiLecturer.put('/test/theme/count', data)
  ),
  getSectionTests: useServiceAction((id: number) =>
    $apiLecturer.get<ITest.GetTest>(`/discipline/section/test/${id}`)
  ),
  getTest: useServiceAction((id: number) =>
    $apiStudent.get<IQuestion.GetTest>(`/test/${id}`)
  ),
  getDisciplineTests: useServiceAction((id: number) =>
    $apiStudent.get<ITest.GetTest>(`/disciplines/test/${id}`)
  ),
  getOpenedTests: useServiceAction(() =>
    $apiStudent.get<ITest.GetTest>('/test/opened')
  ),
  passTest: useServiceAction((data: ITest.PassTest) =>
    $apiStudent.post(`/test/${data.test_id}`, data.answers)
  ),
  testReport: useServiceAction((id: number) =>
    $apiStudent.get(`/test/report/${id}`)
  ),
  getStudentsOpenTest: useServiceAction((id: number) =>
    $apiLecturer.get<ITest.GetStudentsOpenTest>(`/test/students/opened/${id}`)
  ),
  getDoneTestsStudent: useServiceAction(() =>
    $apiStudent.get<ITest.GetTest>('/test/done')
  ),
};
