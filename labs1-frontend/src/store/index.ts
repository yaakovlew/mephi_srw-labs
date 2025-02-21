import { configureStore } from '@reduxjs/toolkit';
import {lab1BApi} from "../api/lab1BApi.ts";
import authReducer from '../reducers/authSlice';
import alternativesReducer from "../reducers/alternativesSlice";
import criteriasReducer from "../reducers/criteriasSlice";
import stepReducer from "../reducers/stepSlice.ts";
import lvl1Reducer from "../reducers/lvl1Slice.ts";
import lvl2Reducer from "../reducers/lvl2Slice.ts";
import evaluationReducer from "../reducers/evaluationSlice.ts";
import qualitativeReducer from "../reducers/qualitativeSlice.ts";
import markReducer from "../reducers/markSlice.ts";

export const store = configureStore({
    reducer: {
        auth: authReducer,
        alternatives: alternativesReducer,
        criterias: criteriasReducer,
        step: stepReducer,
        lvl1: lvl1Reducer,
        lvl2: lvl2Reducer,
        evaluation: evaluationReducer,
        qualitative: qualitativeReducer,
        marks: markReducer,
        [lab1BApi.reducerPath]: lab1BApi.reducer,
    },
    middleware: (getDefaultMiddleware) =>
        getDefaultMiddleware().concat(lab1BApi.middleware),
});

export default store;
