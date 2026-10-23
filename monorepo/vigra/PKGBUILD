# Maintainer: Bruno Pagani <archange@archlinux.org>
# Maintainer: Caleb Maclennan <caleb@alerque.com>
# Contributor: AndyRTR <andyrtr@archlinux.org>
# Contributor: Lukas Jirkovsky <l.jirkovsky@gmail.com>

pkgbase=vigra
pkgname=(vigra python-vigra)
pkgver=1.12.4
pkgrel=3
pkgdesc='Computer vision library'
arch=(x86_64)
url="https://ukoethe.github.io/$pkgname"
_url="https://github.com/ukoethe/$pkgname"
license=(MIT)
makedepends=(
  boost
  boost-libs
  cmake
  fftw
  hdf5
  imath
  libjpeg-turbo
  libpng
  libtiff
  openexr
  python-h5py
  python-matplotlib
  python-numpy
  python-scipy
  zlib
)
checkdepends=(python-pytest)
_pkgver="Version-${pkgver//./-}"
_archive="$pkgname-$_pkgver"
source=("$_url/archive/$_pkgver/$_archive.tar.gz"
        vigra-openexr3.patch)
sha256sums=('fa516eeb19f72ae98a6d5a284c033bc0e0070cb73a8c72985696d2ad73a9292a'
            '2572717e39f916c7c463b8f49306b683c5923906be2721f7da181c58b3a2b34b')

prepare() {
  cd "$_archive"
  patch -p1 -i ../vigra-openexr3.patch

  # Delete invocation of failing test with hdf5 2.0
  # https://github.com/ukoethe/vigra/issues/615
  sed -i '/add( testCase( &ClassifierTest::HDF5ImpexTest));/d' test/classifier/test.cxx
}

build() {
  local cmake_flags=(
    -D CMAKE_INSTALL_PREFIX=/usr
    -D PYTHON_EXECUTABLE=/usr/bin/python
    -D DOCINSTALL=share/doc
    -D WITH_OPENEXR=true
    -D WITH_VIGRANUMPY=1
    # -D BUILD_DOCS=OFF
  )
  cmake -B build -S "$_archive" "${cmake_flags[@]}"
  make -C build
}

check() {
  make -C build -j1 -k check
}

_pick() {
  local p="$1" f d; shift
  for f; do
    d="$srcdir/$p/${f#$pkgdir/}"
    mkdir -p "$(dirname "$d")"
    mv "$f" "$d"
    rmdir -p --ignore-fail-on-non-empty "$(dirname "$f")"
  done
}

package_vigra() {
  depends=(
    libgcc
    libstdc++
    glibc
    hdf5
    imath
    libjpeg-turbo
    libpng
    libtiff
    openexr
    zlib
  )
  optdepends=(
    'python: for vigra-config'
    'python-vigra: for Python bindings'
  )

  make -C build DESTDIR="$pkgdir" install
  install -Dm0644 -t "$pkgdir/usr/share/licenses/$pkgname/" "${_archive}/LICENSE.txt"

  cd "$pkgdir"
  _pick python-$pkgbase usr/lib/python*
  _pick python-$pkgbase usr/lib/vigranumpy
}

package_python-vigra() {
  pkgdesc+=' - Python bindings'
  depends=(
    boost-libs
    fftw
    libgcc
    libstdc++
    glibc
    hdf5
    python
    python-h5py
    python-matplotlib
    python-numpy
    python-scipy
    vigra
  )

  mv -v $pkgname/* "$pkgdir"
  install -Dm0644 -t "$pkgdir/usr/share/licenses/$pkgname/" "$pkgbase-Version-${pkgver//./-}/LICENSE.txt"
}
