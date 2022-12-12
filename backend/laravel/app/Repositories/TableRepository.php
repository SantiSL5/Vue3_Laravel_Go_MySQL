<?php

namespace App\Repositories;
use App\Models\Table;
use App\Interfaces\TableRepositoryInterface;
use App\Http\Requests\Table\CreateTable;
use App\Http\Requests\Table\UpdateTable;
use Illuminate\Database\Eloquent\Builder;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Support\Facades\DB;

class TableRepository implements TableRepositoryInterface
{

    public function getAllTablesRepo()
    {
        return Table::all();
    }

    public function getTableByIdRepo($id) 
    {
        return Table::where('id', $id)->firstOrFail();
    }
    
    public function getTableByCodeRepo($code) 
    {
        Table::where('code', $code)->firstOrFail();
    }

    public function createTableRepo(CreateTable $request)
    {
        return Table::create($request->validated());
    }

    public function updateTableRepo(UpdateTable $request, $id)
    {
        return Table::where('id', $id)->update($request->validated());
    }

    public function deleteTableRepo($id)
    {
        return Table::where('id', $id)->delete();
    }
}
