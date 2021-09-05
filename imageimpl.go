//go:build impl

package faithtop

import (
	"fmt"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type ImageImpl struct {
	*WidgetImpl
	image         *widgets.QGraphicsView
	width, height int
}

func init() {
	newImageImpl = func() IImage {
		v := &ImageImpl{
			image: widgets.NewQGraphicsView(nil),
		}
		v.WidgetImpl = widgetImplFrom(v.image.QWidget_PTR())
		return v
	}
}

func (i *ImageImpl) Assign(v *IImage) IImage {
	*v = i
	return i
}

func (i *ImageImpl) ScaleTo(width, height int) IImage {
	i.width = width
	i.height = height
	return i
}

func (i *ImageImpl) Src(url string) IImage {
	var imageReader = gui.NewQImageReader3(url, core.NewQByteArray())

	scene := widgets.NewQGraphicsScene(nil)

	if imageReader.SupportsAnimation() {
		// instead of reading from file(disk) again, we take from memory
		var movie = gui.NewQMovie3(url, core.NewQByteArray(), nil)
		if i.width > 0 && i.height > 0 {
			movie.SetScaledSize(core.NewQSize2(i.width, i.height))
		}
		// see http://stackoverflow.com/questions/5769766/qt-how-to-show-gifanimated-image-in-qgraphicspixmapitem
		var movieLabel = widgets.NewQLabel(nil, core.Qt__Widget)
		movieLabel.SetMovie(movie)
		movie.ConnectFrameChanged(func(frameNumber int) {
			if frameNumber == movie.FrameCount()-1 {
				movie.Start()
			}
		})
		movie.Start()
		scene.AddWidget(movieLabel, core.Qt__Widget)
	} else {
		var pixmap = gui.QPixmap_FromImageReader(imageReader, core.Qt__AutoColor)
		pixmap = pixmap.Scaled2(i.width, i.height, core.Qt__KeepAspectRatio, core.Qt__FastTransformation)
		item := widgets.NewQGraphicsPixmapItem2(pixmap, nil)
		scene.AddItem(item)
	}

	fmt.Println("Set scene:", url)
	i.image.SetScene(scene)
	return i
}
