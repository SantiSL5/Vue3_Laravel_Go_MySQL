<?php

namespace App\Providers;

use App\Repositories\TableRepository;
use App\Repositories\CategoryRepository;
use App\Repositories\DishTypeRepository;
use App\Repositories\DishRepository;
use App\Repositories\UserRepository;
use App\Interfaces\TableRepositoryInterface;
use App\Interfaces\CategoryRepositoryInterface;
use App\Interfaces\DishTypeRepositoryInterface;
use App\Interfaces\DishRepositoryInterface;
use App\Interfaces\UserRepositoryInterface;
use Illuminate\Support\ServiceProvider;
use Illuminate\Http\Resources\Json\JsonResource;

class AppServiceProvider extends ServiceProvider
{
    /**
     * Register any application services.
     *
     * @return void
     */
    public function register()
    {
        $this->app->bind(CategoryRepositoryInterface::class, function () {
            return new CategoryRepository();
        });

        $this->app->bind(TableRepositoryInterface::class, function () {
            return new TableRepository();
        });

        $this->app->bind(DishTypeRepositoryInterface::class, function () {
            return new DishTypeRepository();
        });

        $this->app->bind(DishRepositoryInterface::class, function () {
            return new DishRepository();
        });
        $this->app->bind(UserRepositoryInterface::class, function () {
            return new UserRepository();
        });
    }

    /**
     * Bootstrap any application services.
     *
     * @return void
     */
    public function boot()
    {
        JsonResource::withoutWrapping();
    }
}
