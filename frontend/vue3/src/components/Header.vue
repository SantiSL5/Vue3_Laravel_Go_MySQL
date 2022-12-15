<template>
    <!-- <nav class="navbar navbar-expand-md bg-dark navbar-dark">
        <span class="navbar-brand">TodoList App</span>
        <button class="navbar-toggler" type="button" @click="changeIsNavShow">
            <span class="navbar-toggler-icon"></span>
        </button>

        <div :class="navClass">
            <ul class="navbar-nav">
                <li class="nav-item">
                    <router-link class="nav-link" to="/home">Home</router-link>
                </li>
                <li class="nav-item">
                    <router-link class="nav-link" to="/shop">Shop</router-link>
                </li>
                <li class="nav-item">
                    <router-link class="nav-link" to="/categories">CRUD Categories</router-link>
                </li>
            </ul>
        </div>
    </nav> -->

    <nav class="navbar navbar-expand-lg navbar-dark bg-dark">
        <div class="container-fluid">
            <a class="navbar-brand" href="#">Namazu</a>
            <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarText"
                aria-controls="navbarText" aria-expanded="false" aria-label="Toggle navigation">
                <span class="navbar-toggler-icon"></span>
            </button>
            <div class="collapse navbar-collapse" id="navbarText">
                <ul class="navbar-nav me-auto mb-2 mb-lg-0">
                    <li class="nav-item">
                        <router-link class="nav-link active" aria-current="page" href="#" to="/home">Home</router-link>
                    </li>
                    <li class="nav-item">
                        <router-link class="nav-link" href="#" to="/shop">Shop</router-link>
                    </li>
                    <li class="nav-item">
                        <router-link class="nav-link" href="#" to="/menus">Menus</router-link>
                    </li>
                </ul>
                <span class="navbar-text">
                    Placeholder
                </span>
            </div>
        </div>
    </nav>



</template>

<script>

export default {
}
</script>

<style>

</style>