<?php

namespace App\Providers;

use App\Repositories\TableRepository;
use App\Repositories\CategoryRepository;
use App\Interfaces\TableRepositoryInterface;
use App\Interfaces\CategoryRepositoryInterface;
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
        $this->app->bind(CategoryRepositoryInterface::class, function (){
            return new CategoryRepository();
        });

        $this->app->bind(TableRepositoryInterface::class, function (){
            return new TableRepository();
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
