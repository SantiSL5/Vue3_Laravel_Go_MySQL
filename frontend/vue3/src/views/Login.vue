<template>
    <div class="d-flex justify-content-center">
        <div class="col-4" id="login" v-if="state.formView == 'login'">
            <form action="/action_page.php">
                <div class="mb-3 mt-3">
                    <label for="loginEmail" class="form-label">Email:</label>
                    <input type="email" class="form-control" id="loginEmail" placeholder="Enter email"
                        name="loginEmail">
                </div>
                <div class="mb-3">
                    <label for="loginPwd" class="form-label">Password:</label>
                    <input type="password" class="form-control" id="loginPwd" placeholder="Enter password"
                        name="loginPwd">
                </div>
                <button type="submit" class="btn btn-primary">Submit</button>
                <div class="mt-2">
                    Don't have an account, <b @click="state.formView = 'register'" class="pointer">create one</b>.
                </div>
            </form>
        </div>

        <div class="col-4" id="register" v-if="state.formView == 'register'">
            <form action="/action_page.php">
                <div class="mb-3 mt-3">
                    <label for="registerUsername" class="form-label">Username:</label>
                    <input class="form-control" id="username" placeholder="Enter username" name="username">
                </div>
                <div class="mb-3 mt-3">
                    <label for="registerEmail" class="form-label">Email:</label>
                    <input type="email" class="form-control" id="registerEmail" placeholder="Enter email" name="registerEmail">
                </div>
                <div class="mb-3">
                    <label for="registerPwd" class="form-label">Password:</label>
                    <input type="password" class="form-control" id="registerPwd" placeholder="Enter password" name="registerPwd">
                </div>
                <div class="mb-3">
                    <label for="registerPwd2" class="form-label">Repeat password:</label>
                    <input type="password" class="form-control" id="registerPwd2" placeholder="Enter password" name="registerPwd2">
                </div>
                <div class="form-check mb-3">
                    <label class="form-check-label">
                        <input class="form-check-input" type="checkbox" name="remember"> Remember me
                    </label>
                </div>
                <button type="submit" class="btn btn-primary">Submit</button>
                <div class="mt-2">
                    Already have an account, <b @click="state.formView = 'login'" class="pointer">log in</b>.
                </div>
            </form>
        </div>
    </div>
</template>

  
<script>
import { reactive } from 'vue';
// import { useVuelidate } from '@vuelidate/core'
// import { required, email } from '@vuelidate/validators'
// import { useStore } from 'vuex';
// import Constant from '../Constant';
export default {
    setup() {
        const state = reactive({
            formView: "login"
        })
        return { state }
    }
}
</script>
  
<style>
.pointer {
    cursor: pointer;
}
</style>
