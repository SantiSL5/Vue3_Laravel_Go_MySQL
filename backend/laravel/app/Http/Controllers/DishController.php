<?php

namespace App\Http\Controllers;

use App\Http\Requests\Dish\CreateDish;
use App\Http\Requests\Dish\UpdateDish;
use App\Http\Resources\DishResource;
use App\Models\Dish;
use App\Models\DishType;
use App\Services\DishService;

use Illuminate\Http\JsonResponse;

class DishController extends Controller
{
    private $dishService;

    public function __construct(DishService $dishTypeService)
    {
        $this->dishService = $dishTypeService;
    }

    /**
     * Display a listing of the resource.
     *
     * @return \Illuminate\Http\Response
     */
    public function index()
    {
        return $this->dishService->getAllDishesService();
    }

    /**
     * Show the form for creating a new resource.
     *
     * @return \Illuminate\Http\Response
     */
    public function create()
    {
        //
    }

    /**
     * Store a newly created resource in storage.
     *
     * @param  \App\Http\Requests\StoreDishRequest  $request
     * @return \Illuminate\Http\Response
     */
    public function store(CreateDish $request)
    {
        return $this->dishService->createDishService($request);
    }

    /**
     * Display the specified resource.
     *
     * @param  \App\Models\Dish  $dish
     * @return \Illuminate\Http\Response
     */
    public function show($id)
    {
        return $this->dishService->getDishByIdService($id);
    }

    /**
     * Show the form for editing the specified resource.
     *
     * @param  \App\Models\Dish  $dish
     * @return \Illuminate\Http\Response
     */
    public function edit(Dish $dish)
    {
        //
    }

    /**
     * Update the specified resource in storage.
     *
     * @param  \App\Http\Requests\UpdateDishRequest  $request
     * @param  \App\Models\Dish  $dish
     * @return \Illuminate\Http\Response
     */
    public function update(UpdateDish $request, $id)
    {
        return $this->dishService->updateDishService($request,$id);
    }

    /**
     * Remove the specified resource from storage.
     *
     * @param  \App\Models\Dish  $dish
     * @return \Illuminate\Http\Response
     */
    public function destroy($id)
    {
        return $this->dishService->deleteDishService($id);
    }
}
