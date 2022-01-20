import { createStore } from "redux";

const defaultState = {}

// Redux Store
const rootReducer = (state = defaultState, action) => {
    //response to action
    switch (action.type) {

        // Return the current state if action doesn't match one this reducer cares about
        default:
            return state
    }
};

export const store = createStore(rootReducer);
