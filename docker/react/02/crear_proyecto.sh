#!/bin/bash  
#Con zsh, da error, read -p, " -p: no coprocess".

read -p "Tag imagen --- id docker hub/nombre imagen:version --- : " tag_imagen
read -p "Confirmar tag imagen --- $tag_imagen --- (S/N) :" confirm && [[ $confirm == [yY] || $confirm == [yY][eE][sS] ||  $confirm == [sS] || $confirm == [sS][iI] ]] || exit 1 
read -p "Nombre container: " nombre_container
read -p "Confirmar nombre container --- $nombre_container --- (S/N) :" confirm && [[ $confirm == [yY] || $confirm == [yY][eE][sS] ||  $confirm == [sS] || $confirm == [sS][iI] ]] || exit 1 
docker build . -f Dockerfile.react.install -t $tag_imagen \
&& docker run --name $nombre_container -v ${PWD}:/app $tag_imagen create-react-app ./proyecto \
&& sudo chown -R pablo:pablo ./proyecto \
&& cp -r ./docker/. ./proyecto