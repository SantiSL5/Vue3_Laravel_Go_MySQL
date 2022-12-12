<?php

namespace App\Interfaces;
use App\Http\Requests\Table\CreateTable;
use App\Http\Requests\Table\UpdateTable;

interface TableRepositoryInterface
{
    function getAllTablesRepo();
    function getTableByIdRepo($id);
    function getTableByCodeRepo($code);
    function createTableRepo(CreateTable $request);
    function updateTableRepo(UpdateTable $request,$id);
    function deleteTableRepo($id);
}
