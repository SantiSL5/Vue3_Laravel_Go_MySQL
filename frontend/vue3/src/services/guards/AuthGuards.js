import store from '../../store';
import Constant from '../../Constant';
import UserServiceAdmin from '../admin/UserService';

export default {

    async authGuardAdmin(to, from, next) {
        try {
            if (localStorage.getItem('tokenAdmin')) {
                const response = await UserServiceAdmin.profile();
                if (response.status === 200) {
                    next();
                }
            } else {
                next('/login');
            }
        } catch (error) {
            store.dispatch(`userClient/${Constant.LOGOUT}`);
        }
    },

    noAuthGuard(to, from, next) {
        if (!localStorage.getItem('token') && !localStorage.getItem('tokenAdmin')) {
            next();
        } else {
            next('/home');
        }
    },
}
