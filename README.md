# Proyecto TFG - Comparación de Tecnologías de API

Este proyecto tiene como objetivo realizar una comparación exhaustiva entre tres tecnologías de API: REST, GraphQL y gRPC. El propósito es evaluar el rendimiento, la escalabilidad y la facilidad de uso de cada una de estas tecnologías, y proporcionar una guía práctica para los desarrolladores que deseen seleccionar la opción más adecuada para sus proyectos.

## Descripción general

Este repositorio contiene el código fuente de una aplicación de ejemplo que implementa las tres tecnologías de API mencionadas. Se han desarrollado tres servidores en Go utilizando las librerías Fiber, graphql-go y grpc-go para los backends REST, GraphQL y gRPC, respectivamente.

Además, se ha creado un cliente único en TypeScript que actúa como interfaz de línea de comandos (CLI). Este cliente se conecta a cada uno de los servidores API mediante las librerías axios, graphql-request y grpc-js, lo que permite a los usuarios interactuar con la aplicación y realizar operaciones CRUD (create, read, update, delete) sobre la base de datos.

Los servidores están conectados a dos versiones de la base de datos Redis: RedisJSON y Redis Core. Estas bases de datos se gestionan y monitorizan manualmente utilizando RedisInsight. La conexión entre los servidores y las bases de datos se realiza mediante el cliente Go Redis.

![architecture](/assets/diagrams/Architecture.png)

## Evaluación de rendimiento y análisis

Para evaluar el rendimiento de cada tecnología de API, se han implementado pruebas de carga utilizando Postman. Además, se ha desarrollado un script en Python utilizando las librerías pandas, NumPy y matplotlib, el cual permite generar informes detallados basados en los resultados de las pruebas. Este análisis proporciona información valiosa sobre el rendimiento y la escalabilidad de cada tecnología.
