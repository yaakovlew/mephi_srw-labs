import { defineStore } from 'pinia';
import { Ref, ref } from 'vue';
import { ISection } from '../models/section/section';
import { SectionService } from '../services/section';

export const useSectionStore = defineStore('sections', () => {
  const sections: Ref<ISection[] | null> = ref(null);

  const getSections = async (id: number) => {
    const res = await SectionService.fetch(id);
    if (res.data) {
      sections.value = res.data.ru;
    }
  };

  const createSection = async (data: ISection.CreateSection) => {
    await SectionService.createSection(data);
  };

  const changeSection = async (data: ISection.Section) => {
    await SectionService.changeSection(data);
  };

  const addLabToSection = async (data: ISection.AddLabToSection) => {
    await SectionService.addLabToSection(data);
  };

  const addTestToSection = async (data: ISection.AddTestToSection) => {
    await SectionService.addTestToSection(data);
  };

  const deleteTestFromSection = async (data: ISection.AddTestToSection) => {
    await SectionService.deleteTestFromSection(data);
  };

  const deleteLabFromSection = async (data: ISection.AddLabToSection) => {
    await SectionService.deleteLabFromSection(data);
  };

  return {
    sections,
    getSections,
    createSection,
    changeSection,
    addLabToSection,
    addTestToSection,
    deleteTestFromSection,
    deleteLabFromSection,
  };
});
