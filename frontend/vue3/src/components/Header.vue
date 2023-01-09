<template>
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
                        <router-link class="nav-link" href="#" to="/reservations">Reservations</router-link>
                    </li>
                    <li class="nav-item">
                        <router-link class="nav-link" href="#" to="/menu">Menu</router-link>
                    </li>
                </ul>
                <span class="navbar-text row">

                    <!-- ////////// -->
                    <div class="col">
                        <router-link class="nav-link" href="#" to="/admin">
                            <button type="button" class="btn btn-info">PanelAdmin</button>
                        </router-link>
                    </div>
                    <!-- ////////// -->
                    <div class="col">
                        <router-link class="nav-link" href="#" to="/login">
                            <button type="button" class="btn btn-info">Login</button>
                        </router-link>
                    </div>
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