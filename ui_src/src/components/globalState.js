import { createStore } from "redux";

const defaultState = {
    lbStatus: 'Active',
}

// Redux Store
const rootReducer = (state = defaultState, action) => {
    //response to action
    switch (action.type) {

        case 'CHANGE_LB_STATUS':
            return {
                ...state,
                lbStatus: action.status,
            }

        // Return the current state if action doesn't match one this reducer cares about
        default:
            return state
    }
};

export const store = createStore(rootReducer);
