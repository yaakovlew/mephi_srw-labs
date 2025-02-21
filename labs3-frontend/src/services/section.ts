import { $apiLecturer } from 'src/boot/axios';
import { useServiceAction } from 'src/utils/service/action';
import { ISection } from '../models/section/section';

export const SectionService = {
  fetch: useServiceAction((id: number) =>
    $apiLecturer.get<ISection.GetSections>(`/discipline/section/${id}`)
  ),
  createSection: useServiceAction((data: ISection.CreateSection) =>
    $apiLecturer.post<ISection.CreateSection>('/discipline/section', data)
  ),
  changeSection: useServiceAction((data: ISection.Section) =>
    $apiLecturer.put<ISection.Section>('/discipline/section', data)
  ),
  addLabToSection: useServiceAction((data: ISection.AddLabToSection) =>
    $apiLecturer.post<ISection.AddLabToSection>(
      '/discipline/section/labaratory',
      data
    )
  ),
  addTestToSection: useServiceAction((data: ISection.AddTestToSection) =>
    $apiLecturer.post<ISection.AddTestToSection>(
      '/discipline/section/test',
      data
    )
  ),
  deleteTestFromSection: useServiceAction((data: ISection.AddTestToSection) =>
    $apiLecturer.delete(
      `/discipline/section/test?section_id=${data.section_id}&test_id=${data.test_id}`
    )
  ),
  deleteLabFromSection: useServiceAction((data: ISection.AddLabToSection) =>
    $apiLecturer.delete(
      `/discipline/section/labaratory?section_id=${data.section_id}&labaratory_id=${data.labaratory_id}`
    )
  ),
};
