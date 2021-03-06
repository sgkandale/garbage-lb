import { createStore } from "redux";

// const dummyServerLoad = {
//     ram: {
//         total: '1024 M',
//         inUse: '996 M',
//         available: '28 M',
//         inUsePercentage: '97.2 %'
//     },
//     storage: {
//         total: '50 G',
//         inUse: '21 G',
//         available: '29 G',
//         inUsePercentage: '42 %',
//     },
//     network: {
//         total: '3 KB/s',
//         up: '2 KB/s',
//         down: '1 KB/s',
//     },
//     cpu: {
//         total: '2.7 GHz',
//         inUse: '1.16 GHz',
//         available: '1.54 GHz',
//         inUsePercentage: '23.2 %',
//     }
// }

// const dummyListeners = [
//     {
//         name: "Some Name",
//         port: "8080",
//         type: 'HTTP',
//         id: "abcd",
//         listening: true,
//         filters: [
//             {
//                 name: "Some Filter",
//                 id: "abcd",
//                 rules: [
//                     {
//                         name: "Some Rule",
//                         id: "abcd",
//                         type: "Path",
//                         value: "/some/path",
//                         subValue: "",
//                         action: "Allow",
//                         enabled: true,
//                     },
//                 ],
//             }
//         ]
//     }
// ]
// const dummyClusters = [
//     {
//         name: "Some Name",
//         id: "abcd",
//         type: 'Logical',
//         policy: 'RoundRobin',
//         timeout: '30',
//         endpoints: [
//             {
//                 id: "abcd",
//                 name: 'Localhost',
//                 address: "0.0.0.0",
//                 port: "8080",
//                 health: "Healthy",
//             }
//         ],
//         health: {
//             status: 'Healthy',
//             healthyCount: 1,
//             unhealthyCount: 0,
//             degradedCount: 0,
//         },
//     }
// ]

const defaultState = {
    lbStatus: 'Unknown',
    listeners: [],
    clusters: [],
    serverLoad: {
        loading: true,
    },
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

        case 'SET_SERVER_LOAD':
            return {
                ...state,
                serverLoad: action.serverLoad,
            }

        case 'SET_SERVER_STATUS':
            return {
                ...state,
                lbStatus: action.lbStatus,
            }

        case 'SET_LISTENERS':
            return {
                ...state,
                listeners: action.listeners,
            }

        case 'SET_CLUSTERS':
            return {
                ...state,
                clusters: action.clusters,
            }

        // Return the current state if action doesn't match one this reducer cares about
        default:
            return state
    }
};

export const store = createStore(rootReducer);
