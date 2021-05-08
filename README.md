# DataloggerDemo

<h4>Visual Studio code launch.json</h4>
<p>
{
    // Use IntelliSense para saber los atributos posibles.
    // Mantenga el puntero para ver las descripciones de los existentes atributos.
    // Para más información, visite: https://go.microsoft.com/fwlink/?linkid=830387
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch Package",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${workspaceFolder}/src",
            "env": {
                "MYSQL_HOST": "localhost",
                "MYSQL_PORT":"3306",
                "MYSQL_DATABASE":"datalogger",
                "MYSQL_USER":"api_datalogger",
                "MYSQL_PASSWORD":"12345",
                "BASE_URL":"/dlm/api/v1",
                "PORT":"8367"
            }
        }
    ]
}
</p>
