# fibertel-stats

Hice este programa simple para almacenar en un archivo de texto los niveles de Tx, Rx y SNR de mi conexión Fibertel, útiles al tratar de diagnosticar problemas y al realizar un reclamo. Al ejecutar el mismo se obtienen los niveles [desde los servers de Fibertel](http://provisioning.fibertel.com.ar/asp/nivelesPrima.asp) automáticamente cada 1 minuto y se almacenan en un archivo llamado `fibertel-stats.log`. Llegado el caso de que el archivo exceda los 50Mb de tamaño, se crea automáticamente un nuevo archivo.

Cada medición es almacenada en el archivo `.log` con el siguiente formato:

```
yyyy/mm/dd hh:mm:ss Estado [Tx FreqTx Rx FreqRx MER]
```

`Tx` y `Rx` están en `dBmV`, `FreqTx` y `FreqRx` están en `MHz`, y `MER` está en `dB`.

El estado se representa mediante `✓` si pudo obtener la información y mediante `✕` si hubo un error al hacerlo; este estado también se puede observar en la ventana donde se ejecuta el programa. En el caso de que la obtención de datos falle, se indica la causa en lugar de los valores. Es importante destacar que el estado habla únicamente de la obtención de datos, un estado `✓` **no indica que los valores sean apropiados o que la conexión sea buena**, sino únicamente que se pudieron obtener correctamente.

Este es un ejemplo de la salida del programa, donde podemos ver dos mediciones correctas y dos fallidas por distinta causa:

```
2018/07/01 23:27:18 ✓ [43.5 36600 7 771000 38.7]
2018/07/01 23:27:28 ✓ [43.5 36600 7.1 771000 39.1]
2018/07/01 23:27:39 ✕ Get http://provisioning.fibertel.com.ar/asp/nivelesPrima.asp: net/http: request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)
2018/07/01 23:27:49 ✕ Get http://provisioning.fibertel.com.ar/asp/nivelesPrima.asp: dial tcp: lookup provisioning.fibertel.com.ar: no such host
```

La versión más reciente puede descargarse en [https://github.com/fcingolani/fibertel-stats/releases]().

## ¿Están bien mis valores?

Los valores apropiados varían de acuerdo a la versión de DOCSIS (un estandar internacional para la transmisión de datos), aunque podemos encontrar algunos valores de referencia en [http://www.dslreports.com/faq/16085](), que se detallan a continuación.

### Rx

| Estado                  | Mínimo      | Máximo      |
|-------------------------|-------------|-------------|
| Recomendado             | -7 dBmV     | +7 dBmV     |
| Aceptable               | -10 dBmV    | +10 dBmV    |
| Máximo tolerable        | -15 dBmV    | +15 dBmV    |
| Fuera de Especificación | < -15 dBmV  | > +15 dBmV  |

### Tx

| Versión     | Mínimo    | Máximo    |
|-------------|-----------|-----------|
| DOCSIS 3.1  | +35 dBmV  | +47 dBmV  |
| DOCSIS 3.0  | +35 dBmV  | +49 dBmV  |