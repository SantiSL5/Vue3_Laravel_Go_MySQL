<template>
    <div class="d-flex justify-content-center">
        <div class="col-4" id="login" v-if="state.formView == 'login'">
            <form @submit.prevent="handleSubmit">
                <div class="mb-3 mt-3">
                    <label for="loginEmail" class="form-label">Email:</label>
                    <input type="email" class="form-control" id="loginEmail" placeholder="Enter email" name="loginEmail"
                        v-model="state.loginForm.email">
                    <span class="text-danger">{{ state.errLogin.email }}</span>
                </div>
                <div class="mb-3">
                    <label for="loginPwd" class="form-label">Password:</label>
                    <input type="password" class="form-control" id="loginPwd" placeholder="Enter password"
                        name="loginPwd" v-model="state.loginForm.password">
                    <span class="text-danger">{{ state.errLogin.password }}</span>
                </div>
                <button type="button" class="btn btn-primary" @click="login">Submit</button>
                <div class="mt-2">
                    Don't have an account, <b @click="state.formView = 'register'" class="pointer">create one</b>.
                </div>
            </form>
        </div>

        <div class="col-4" id="register" v-if="state.formView == 'register'">
            <form>
                <div class="mb-3 mt-3">
                    <label for="registerUsername" class="form-label">Username:</label>
                    <input class="form-control" id="username" placeholder="Enter username" name="username"
                        v-model="state.registerForm.username">
                    <span class="text-danger">{{ state.errRegister.username }}</span>
                </div>
                <div class="mb-3 mt-3">
                    <label for="registerEmail" class="form-label">Email:</label>
                    <input type="email" class="form-control" id="registerEmail" placeholder="Enter email"
                        name="registerEmail" v-model="state.registerForm.email">
                    <span class="text-danger">{{ state.errRegister.email }}</span>
                </div>
                <div class="mb-3">
                    <label for="registerPwd" class="form-label">Password:</label>
                    <input type="password" class="form-control" id="registerPwd" placeholder="Enter password"
                        name="registerPwd" v-model="state.registerForm.password">
                    <span class="text-danger">{{ state.errRegister.password }}</span>
                </div>
                <div class="mb-3">
                    <label for="registerPwd2" class="form-label">Repeat password:</label>
                    <input type="password" class="form-control" id="registerPwd2" placeholder="Enter password"
                        name="registerPwd2" v-model="state.registerForm.password2">
                    <span class="text-danger">{{ state.errRegister.password2 }}</span>
                </div>
                <button type="button" class="btn btn-primary" @click="register">Submit</button>
                <div class="mt-2">
                    Already have an account, <b @click="state.formView = 'login'" class="pointer">log in</b>.
                </div>
            </form>
        </div>
    </div>
</template>

  
<script>
import Constant from "../Constant";
import { reactive } from 'vue';
import { useVuelidate } from '@vuelidate/core'
import { required, email, minLength, maxLength } from '@vuelidate/validators'
// import { useStore } from 'vuex';
// import Constant from '../Constant';
import { useStore } from "vuex";


export default {
    setup() {
        const store = useStore();

        const state = reactive({
            formView: "login",
            loginForm: { email: "", password: "" },
            registerForm: { username: "", email: "", password: "", password2: "" },
            errLogin: { email: "", password: "" },
            errRegister: { username: "", email: "", password: "", password2: "" },
        })

        const validationsLogin = {
            email: { required, email },
            password: { required, minLength: minLength(6), maxLength: maxLength(30) }
        }

        const validationsRegister = {
            username: { required },
            email: { required, email },
            password: { required, minLength: minLength(6), maxLength: maxLength(30) },
            password2: { required, minLength: minLength(6), maxLength: maxLength(30) },
        }

        state.loginErr = useVuelidate(validationsLogin, state.loginForm);
        state.registerErr = useVuelidate(validationsRegister, state.registerForm);

        const login = () => {
            if (state.loginErr.email.$invalid == true) {
                if (state.loginErr.email.$model != "") {
                    state.errLogin.email = "Email is invalid";
                } else {
                    state.errLogin.email = "Email is required";
                }
            } else {
                state.errLogin.email = "";
            }
            if (state.loginErr.password.$invalid == true) {
                if (state.loginErr.password.$model.length < 6 && state.loginErr.password.$model.length != 0) {
                    state.errLogin.password = "Min 6 characters";
                } else if (state.loginErr.password.$model.length > 30) {
                    state.errLogin.password = "Max 30 characters";
                } else {
                    state.errLogin.password = "Password is required (6-30)"
                }
            } else {
                state.errLogin.password = "";
            }
            if (state.loginErr.email.$invalid != true && state.loginErr.password.$invalid != true) {
                store.dispatch("userClient/" + Constant.LOGIN, state.loginForm);
            }
        }

        const register = () => {
            if (state.registerErr.email.$invalid == true) {
                if (state.registerErr.email.$model != "") {
                    state.errRegister.email = "Email is invalid";
                } else {
                    state.errRegister.email = "Email is required";
                }
            } else {
                state.errRegister.email = "";
            }

            if (state.registerErr.username.$invalid == true) {
                state.errRegister.username = "Username required";
            } else {
                state.errRegister.username = "";
            }

            if (state.registerErr.password.$invalid == true) {
                if (state.registerErr.password.$model.length < 6 && state.registerErr.password.$model.length != 0) {
                    state.errRegister.password = "Min 6 characters";
                } else if (state.registerErr.password.$model.length > 30) {
                    state.errRegister.password = "Max 30 characters";
                } else {
                    state.errRegister.password = "Password is required (6-30)"
                }
            } else {
                state.errRegister.password = "";
            }

            if (state.registerErr.password.$model !== state.registerErr.password2.$model) {
                state.errRegister.password2 = "Passwords must be equal";
            } else {
                state.errRegister.password2 = "";
            }

            if (state.registerErr.email.$invalid != true && state.registerErr.username.$invalid != true && state.registerErr.password.$invalid != true && state.registerErr.password2.$invalid != true && state.registerErr.password.$model === state.registerErr.password2.$model) {
                delete state.registerForm.password2;

                store.dispatch("userClient/" + Constant.REGISTER, state.registerForm);

                // if (computed(() => store.getters["userClient/getUser"].name != "")) {
                //     router.push({ name: 'home' })
                // }
            }
        }


        return { state, login, register }
    },
}
</script>
  
<style>
.pointer {
    cursor: pointer;
}
</style>
