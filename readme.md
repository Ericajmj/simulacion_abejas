## Simulación de abejas


### Abejas
Las abejas se representan como pares ordenados de coordenadas en metros, por lo tanto son puntos sin área.

Se dibujan con imágenes de 20px.

Tienen una velocidad en metros por segundo y un ángulo en el que se mueven (0º=derecha, 90º=arriba, 180º=izquierda, 270º=abajo).

Las abejas recorren entre 1000m y 3000m.


### Escala

El input de la simulación es una imagen de un mapa del que se debe conocer la escala pixeles a metros.

En el ejemplo actual, se usa una escala de 100m = 50px, por lo que cada 1px = 2.0m.

Utilizar la constante metrosPorPixel para cambiar la escala.#   s i m u l a c i o n _ a b e j a s  
 #   a b e j a s _ z a p o p a n  
 