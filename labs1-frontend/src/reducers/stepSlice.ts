import { createSlice, PayloadAction } from '@reduxjs/toolkit';

const localStorageNameActiveStep = "activeStep1b"
const localStorageNameCompletedSteps = "completedSteps1b"

interface StepState {
    activeStep: number;
    completedSteps: number[];
}

const initialState: StepState = {
    activeStep: JSON.parse(localStorage.getItem(localStorageNameActiveStep) || "0"),
    completedSteps: JSON.parse(localStorage.getItem(localStorageNameCompletedSteps) || "[]"),
};

const stepSlice = createSlice({
    name: 'step',
    initialState,
    reducers: {
        setActiveStep(state, action: PayloadAction<number>) {
            state.activeStep = action.payload;
            localStorage.setItem(localStorageNameActiveStep, JSON.stringify(action.payload)); // Сохраняем в локальное хранилище
        },
        addCompletedStep(state, action: PayloadAction<number>) {
            const stepToComplete = action.payload;
            if (!state.completedSteps.includes(stepToComplete)) {
                state.completedSteps.push(stepToComplete);
            }
            localStorage.setItem(localStorageNameCompletedSteps, JSON.stringify(state.completedSteps));
        }
    },
});

export const getActiveStep = ((state: any) => state.step.activeStep)
export const getCompletedSteps = ((state: any) => state.step.completedSteps)

export const {
    setActiveStep,
    addCompletedStep
} = stepSlice.actions;

export default stepSlice.reducer;
