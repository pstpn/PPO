import EmployeeService from "../services/employee.service";

export const employee = {
    namespaced: true,
    state: {
        message: "",
        profile: null,
        photoURL: "//ssl.gstatic.com/accounts/ui/avatar_2x.png",
        status: {
            filled: false,
        }
    },
    actions: {
        fillProfile({ commit }, formData) {
            return EmployeeService.fillProfile(formData).then(
                response => {
                    commit('setFilled', true);
                    return Promise.resolve(response.data);
                },
                error => {
                    return Promise.reject(error);
                })
        },
        getProfile({ commit }) {
            return EmployeeService.getProfile().then(
                profile => {
                    commit('setFilled', true);
                    commit('setProfile', profile);
                    return Promise.resolve(profile);
                },
                error => {
                    commit('setFilled', false);
                    commit('setProfile', null);
                    return Promise.reject(error);
                })
        },
        getEmployeePhoto({ commit }) {
            return EmployeeService.getEmployeePhoto().then(
                imageURL => {
                    commit('setPhotoURL', imageURL);
                    return Promise.resolve(imageURL);
                },
                error => {
                    return Promise.reject(error);
                }
            )
        },
        getEmployees({ commit }, { searchQuery, searchBy, sortDirection }) {
            return EmployeeService.getEmployees(searchQuery, searchBy, sortDirection).then(
                employees => {
                    return Promise.resolve(employees);
                },
                error => {
                    return Promise.reject(error);
                }
            );
        },
        getEmployee({ commit }, infoCardID) {
            return EmployeeService.getEmployee(infoCardID).then(
                employee => {
                    return Promise.resolve(employee);
                },
                error => {
                    return Promise.reject(error);
                }
            );
        },
        getEmployeeInfoCardPhoto({ commit }, infoCardID) {
            return EmployeeService.getEmployeeInfoCardPhoto(infoCardID).then(
                imageURL => {
                    return Promise.resolve(imageURL);
                },
                error => {
                    return Promise.reject(error);
                }
            )
        },
        confirmEmployeeCard({ commit }, infoCardID) {
            return EmployeeService.confirmEmployeeCard(infoCardID).then(
                response => {
                    return Promise.resolve(response);
                },
                error => {
                    return Promise.reject(error);
                }
            );
        },
        createEmployeePassage({ commit }, passageInfo) {
            return EmployeeService.createEmployeePassage(passageInfo).then(
                response => {
                    return Promise.resolve(response);
                },
                error => {
                    return Promise.reject(error);
                }
            )
        },
    },
    mutations: {
        setFilled(state, filled) {
            state.status.filled = filled;
        },
        setProfile(state, profile) {
            state.status.filled = true;
            state.profile = profile;
            state.photoURL = "//ssl.gstatic.com/accounts/ui/avatar_2x.png";
        },
        setPhotoURL(state, photoURL) {
            state.photoURL = photoURL;
        },
    },
};
