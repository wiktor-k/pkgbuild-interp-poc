# Maintainer: Balló György <ballogyor+arch at gmail dot com>

pkgname=gaphor
pkgver=3.3.2
pkgrel=1
pkgdesc='Simple UML and SysML modeling tool'
arch=(any)
url='https://apps.gnome.org/Gaphor/'
license=(Apache-2.0)
groups=(gnome-circle)
depends=(
  dconf
  gdk-pixbuf2
  glib2
  gtk4
  gtksourceview5
  hicolor-icon-theme
  libadwaita
  pango
  python
  python-babel
  python-better-exceptions
  python-cairo
  python-defusedxml
  python-dulwich
  python-gaphas
  python-generic
  python-gobject
  python-jedi
  python-pillow
  python-pydot
  python-tinycss2
)
makedepends=(
  git
  python-build
  python-installer
  python-poetry-core
  python-sphinx-intl
  python-wheel
)
source=("git+https://github.com/gaphor/gaphor.git#tag=$pkgver")
b2sums=('9677c539dd9a3ba4930a14c53a108f72b92009952587b0e309582fe2dfa6ff2f667beb2ec9e96aaf88368b06a594f491aab4a9e0a8c402000ff2fac621d540d8')

build() {
  cd $pkgname
  python scripts/build_babel.py
  python -m build --wheel --no-isolation
}

package() {
  cd $pkgname
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -Dm644 -t "$pkgdir/usr/share/glib-2.0/schemas/" gaphor/ui/installschemas/org.gaphor.Gaphor.gschema.xml
  cp -r gaphor/locale "$pkgdir"/usr/lib/python*/site-packages/gaphor/

  install -Dm644 data/org.gaphor.Gaphor.appdata.xml "$pkgdir/usr/share/metainfo/org.gaphor.Gaphor.appdata.xml"
  install -Dm644 data/org.gaphor.Gaphor.desktop "$pkgdir/usr/share/applications/org.gaphor.Gaphor.desktop"
  install -Dm644 data/org.gaphor.Gaphor.service "$pkgdir/usr/share/dbus-1/services/org.gaphor.Gaphor.service"
  install -Dm644 data/org.gaphor.Gaphor.xml "$pkgdir/usr/share/mime/packages/org.gaphor.Gaphor.xml"
  install -Dm644 data/logos/gaphor-24x24.png "$pkgdir/usr/share/icons/hicolor/24x24/apps/org.gaphor.Gaphor.png"
  install -Dm644 data/logos/gaphor-48x48.png "$pkgdir/usr/share/icons/hicolor/48x48/apps/org.gaphor.Gaphor.png"
  install -Dm644 data/logos/org.gaphor.Gaphor.svg "$pkgdir/usr/share/icons/hicolor/scalable/apps/org.gaphor.Gaphor.svg"
  install -Dm644 data/logos/org.gaphor.Gaphor-symbolic.svg "$pkgdir/usr/share/icons/hicolor/symbolic/apps/org.gaphor.Gaphor-symbolic.svg"
}
