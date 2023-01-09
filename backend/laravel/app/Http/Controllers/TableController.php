<?php

namespace App\Http\Controllers;

use App\Http\Requests\Table\CreateTable;
use App\Http\Requests\Table\UpdateTable;
use App\Http\Resources\TableResource;
use App\Models\Table;
use App\Models\Category;
use App\Services\TableService;

use Illuminate\Http\JsonResponse;

class TableController extends Controller
{
    private $tableService;

    public function __construct(TableService $tableService)
    {
        $this->tableService = $tableService;
    }

    /**
     * Display a listing of the resource.
     *
     * @return \Illuminate\Http\Response
     */
    public function index()
    {
        return $this->tableService->getAllTablesService();
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
     * @param  \App\Http\Requests\StoreTableRequest  $request
     * @return \Illuminate\Http\Response
     */
    public function store(CreateTable $request)
    {
        return $this->tableService->createTableService($request);
    }

    /**
     * Display the specified resource.
     *
     * @param  \App\Models\Table  $table
     * @return \Illuminate\Http\Response
     */
    public function show($id)
    {
        return $this->tableService->getTableByIdService($id);
    }

    /**
     * Show the form for editing the specified resource.
     *
     * @param  \App\Models\Table  $table
     * @return \Illuminate\Http\Response
     */
    public function edit(Table $table)
    {
        //
    }

    /**
     * Update the specified resource in storage.
     *
     * @param  \App\Http\Requests\UpdateTableRequest  $request
     * @param  \App\Models\Table  $table
     * @return \Illuminate\Http\Response
     */
    public function update(UpdateTable $request, $id)
    {
        return $this->tableService->updateTableService($request,$id);
    }

    /**
     * Remove the specified resource from storage.
     *
     * @param  \App\Models\Table  $table
     * @return \Illuminate\Http\Response
     */
    public function destroy($id)
    {
        return $this->tableService->deleteTableService($id);
    }
}
