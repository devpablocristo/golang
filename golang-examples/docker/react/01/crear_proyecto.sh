#!/bin/zsh
docker build . -f Dockerfile.react.install -t temp_react_builder_img_1fnnyu78923drscoi \
&& docker run --name temp_react_builder_container_1fnnyu78923drscoi -v ${PWD}:/app temp_react_builder_img_1fnnyu78923drscoi create-react-app ./temp_proyecto_react \
&& docker rm temp_react_builder_container_1fnnyu78923drscoi \
&& docker rmi temp_react_builder_img_1fnnyu78923drscoi \
&& sudo chown -R pablo:pablo ./temp_proyecto_react \
&& mkdir ./docker_react_install \
&& find . -maxdepth 1 -type f -print0 | xargs -0 mv -t ./docker_react_install \
&& mv ./temp_proyecto_react/*(DN) . \
&& mv ./temp_dockerfiles/*(DN) . \
&& rm -r ./temp_proyecto_react \
&& rm -r ./temp_dockerfiles

#podria hacerlo mas sofiticado con estos otros comandos, pero asi como esta funcina perfecto. KISS
#&& echo $1 | sudo -S chown -R pablo:pablo ./temp_proyecto_react \
#read -p "Enter fullname: " fullname