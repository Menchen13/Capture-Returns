#####
# This is a working example of setting up tesseract/gosseract,
# and also works as an example runtime to use gosseract package.
# You can just hit `docker run -it --rm otiai10/gosseract`
# to try and check it out!
#####
FROM gocv/opencv:latest

RUN apt-get update -qq

# You need librariy files and headers of tesseract and leptonica.
# When you miss these or LD_LIBRARY_PATH is not set to them,
# you would face an error: "tesseract/baseapi.h: No such file or directory"
RUN apt-get install -y -qq libtesseract-dev libleptonica-dev
#sudo for dev container
RUN apt-get install -y -qq sudo

# In case you face TESSDATA_PREFIX error, you minght need to set env vars
# to specify the directory where "tessdata" is located.
ENV TESSDATA_PREFIX=/usr/share/tesseract-ocr/5/tessdata/

# Load languages.
# These {lang}.traineddata would b located under ${TESSDATA_PREFIX}/tessdata.
RUN apt-get install -y -qq \
  tesseract-ocr-eng \
  tesseract-ocr-deu \
  tesseract-ocr-jpn
# See https://github.com/tesseract-ocr/tessdata for the list of available languages.
# If you want to download these traineddata via `wget`, don't forget to locate
# downloaded traineddata under ${TESSDATA_PREFIX}/tessdata.

# Setup your cool project with go.mod.
#WORKDIR /gocv

#RUN git clone https://github.com/hybridgroup/gocv.git && cd gocv && make install

#all my project files will be added using a volume during development time
#I also shouldnt have to get any dependencies, as they should be installed acording to the go.mod and go.sum file
#which are passed with the volume

RUN useradd -ms /bin/bash dev

# Allow the "dev" user to run sudo without a password
RUN echo "dev ALL=(ALL) NOPASSWD: ALL" >> /etc/sudoers


USER dev


WORKDIR /workdir 
# Let's have gosseract in your project and test it.
#RUN go get -t github.com/otiai10/gosseract/v2

# Now, you've got complete environment to play with "gosseract"!
# For other OS, check https://github.com/otiai10/gosseract/tree/main/test/runtimes

# Try `docker run -it --rm otiai10/gosseract` to test this environment.
#CMD go test -v github.com/otiai10/gosseract/v2
CMD go test -v ./...
 