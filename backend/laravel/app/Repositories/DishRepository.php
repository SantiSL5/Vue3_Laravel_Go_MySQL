<?php

namespace App\Repositories;
use App\Models\Dish;
use App\Interfaces\DishRepositoryInterface;
use App\Http\Requests\Dish\CreateDish;
use App\Http\Requests\Dish\UpdateDish;
use Illuminate\Database\Eloquent\Builder;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Support\Facades\DB;

class DishRepository implements DishRepositoryInterface
{

    public function getAllDishesRepo()
    {
        return Dish::all();
    }

    public function getDishByIdRepo($id) 
    {
        return Dish::where('id', $id)->firstOrFail();
    }
    
    public function getDishByCodeRepo($code) 
    {
        Dish::where('code', $code)->firstOrFail();
    }

    public function createDishRepo(CreateDish $request)
    {
        return Dish::create($request->validated());
    }

    public function updateDishRepo(UpdateDish $request, $id)
    {
        return Dish::where('id', $id)->update($request->validated());
    }

    public function deleteDishRepo($id)
    {
        return Dish::where('id', $id)->delete();
    }
}
