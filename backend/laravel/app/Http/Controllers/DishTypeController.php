<?php

namespace App\Http\Controllers;

use App\Http\Requests\DishType\CreateDishType;
use App\Http\Requests\DishType\UpdateDishType;
use App\Models\DishType;
use App\Services\DishTypeService;

use Illuminate\Http\JsonResponse;

class DishTypeController extends Controller
{
    private $dishTypeService;

    public function __construct(DishTypeService $dishTypeService)
    {
        $this->dishTypeService = $dishTypeService;
    }

    /**
     * Display a listing of the resource.
     *
     * @return \Illuminate\Http\Response
     */
    public function index()
    {
        return $this->dishTypeService->getAllDishTypesService();
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
     * @param  \App\Http\Requests\DishType\CreateDishType  $request
     * @return \Illuminate\Http\Response
     */
    public function store(CreateDishType $request)
    {
        return $this->dishTypeService->createDishTypeService($request);
    }

    /**
     * Display the specified resource.
     *
     * @param  \App\Models\DishType  $dishType
     * @return \Illuminate\Http\Response
     */
    public function show($id)
    {
        return $this->dishTypeService->getDishTypeByIdService($id);
    }

    /**
     * Show the form for editing the specified resource.
     *
     * @param  \App\Models\DishType  $dishType
     * @return \Illuminate\Http\Response
     */
    public function edit(DishType $dishType)
    {
        //
    }

    /**
     * Update the specified resource in storage.
     *
     * @param  \App\Http\Requests\DishType\UpdateDishType $request
     * @param  \App\Models\DishType  $dishType
     * @return \Illuminate\Http\Response
     */
    public function update(UpdateDishType $request, $id)
    {
        return $this->dishTypeService->updateDishTypeService($request,$id);
    }

    /**
     * Remove the specified resource from storage.
     *
     * @param  \App\Models\DishType  $dishType
     * @return \Illuminate\Http\Response
     */
    public function destroy($id)
    {
        return $this->dishTypeService->deleteDishTypeService($id);
    }
}
