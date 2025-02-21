import { createSlice, PayloadAction } from "@reduxjs/toolkit";

const localStorageName = "alternatives"


interface AlternativesState {
    alternatives: string[];
}

const initialState: AlternativesState = {
    alternatives: JSON.parse(localStorage.getItem(localStorageName) || "[]"), // Инициализация из localStorage
};

const alternativesSlice = createSlice({
    name: "alternatives",
    initialState,
    reducers: {
        setAlternatives: (state, action: PayloadAction<string[]>) => {
            state.alternatives = action.payload;
            localStorage.setItem(localStorageName, JSON.stringify(state.alternatives));
        },
    },
});

export const getAlternativesNumber = (state: AlternativesState) => state.alternatives.alternatives.length;
export const getAlternatives = (state: AlternativesState) => state.alternatives.alternatives;

export const {
    setAlternatives,
} = alternativesSlice.actions;
export default alternativesSlice.reducer;
