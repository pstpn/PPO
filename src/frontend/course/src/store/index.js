import { createStore } from "vuex";
import { auth } from "./auth.module";
import { employee } from "./employee.module";

const store = createStore({
    modules: {
        auth,
        employee
    },
});

export default store;