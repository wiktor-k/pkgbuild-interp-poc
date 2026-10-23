# Maintainer: George Rawlinson <grawlinson@archlinux.org>

pkgname=python-curl_cffi
pkgver=0.16.3
_curl_version=curl-8_21_0
_curl_impersonate_version=2.2.2
pkgrel=1
pkgdesc='Python FFI binding for curl-impersonate'
arch=(x86_64)
url='https://github.com/lexiforest/curl_cffi'
license=(MIT)
depends=(
  glibc
	"curl-impersonate=${_curl_impersonate_version}"
	python
	python-certifi
	python-cffi
	python-eventlet
	python-gevent
	python-typing_extensions
  python-rich
)
makedepends=(
  git
	python-build
	python-installer
	python-setuptools
	python-wheel
	python-python-multipart
	unzip
)
optdepends=(
  'python-orjson: speed and memory optimized JSON parsing'
  'python-markdownify'
  'python-lxml-html-clean'
  'python-readability-lxml'
)
source=(
	"$pkgname::git+$url#tag=v$pkgver"
  "${_curl_version}.tar.gz::https://github.com/curl/curl/archive/${_curl_version}.tar.gz"
  "curl-impersonate::git+https://github.com/lexiforest/curl-impersonate#tag=v${_curl_impersonate_version}"
	use-system-curl-impersonate.patch
  no-download.patch
)
sha512sums=('64b23be54e4e7a809f6e5756157e057770240b2fb4c377740a20dfd8c39d79131b2c97dbdf536cb98827dd7a4318374837efd3499aff6416ca7c01b64f468f17'
            '0ab6c99c3d5b86fb65c526db517c3159b11db2f8d82552d635c4887059c0602288603c93b754ce0ec543ea2f275122ccec2c8dcd866c2611b5b949c728ee72df'
            'c59e564a07357f6f2b67eb961f5521c16bf7aa3e47e6813265b6707984fe821793f199511855f1afa2f1224aea628a7d6315366c1c1d840487959bf74ddf672d'
            'f543f6aaf45809bcc144168db29cdf5cd28e21ccf6c8efb535e241111606a66115d674dc62540a4b2b17e96ce1de6e5c2696eff490424643235f68c633499700'
            'f9f74740d574cfea9d3c8f6f76c07359f896bed30aaf3c4791a4c92189d365ce61d50f60666d7cf2f25f3f4b20e90efe7f8176ed654870f46d20cc566834243e')
b2sums=('6c5dc04db0420bb6bcc5ef6633d4325043310e40759cb142264d478720b70d94316a6d419d327844baa8c26ad98017462d7f2280504247f178f0ace78badf4a0'
        '3e4df97e213a5aa56ef144cdfd085ea51b2902cf09d6006315e8036156eded548814a43983cd2eea9a653d635219675a8223682f2c98ce4379a8d6f5cd14f63c'
        '69cbe84a8cd95036777458639cbec42d9613ef10d83ed53c21c8fc1d5dfdacc8c3920efd8cddae56f1ba96908ab0356a9104bc1276260bc5a668f9fd122c194a'
        '3a83f7609c62416237ae8d8bd6a894e7d288070662a9dbfb7e776aa2035794ae94dcf09b4701e5700510fdcb51f11478ea361a74aa24ecf383b4dd530d0e89d0'
        '01235672b138af083d1f3d292e9c7b78bc94f38934c87e532b40bd79bf10b5a9b791f91d34aa9c84ada6c73b42002e1d44381eaa7b41b7fee64219410c03bf17')

prepare() {
  cd "$pkgname"

  # use system curl-impersonate
  patch -p1 -i "$srcdir/use-system-curl-impersonate.patch"

  # shuffle around dependencies to match folder structure expected in Makefile
  mv "$srcdir/curl-${_curl_version}" "${_curl_version}"
  mv "$srcdir/curl-impersonate" "curl-impersonate-${_curl_impersonate_version}"

  # patch Makefile for file verification
  patch -p1 -i "$srcdir/no-download.patch"
  sed \
    -e "s/@CURL_IMPERSONATE_VERSION@/${_curl_impersonate_version}/" \
    -e "s/@CURL_VERSION@/${_curl_version}/" \
    -i Makefile \
    -i scripts/build.py

  make preprocess
}

build() {
  cd "$pkgname"

  python -m build --wheel --no-isolation
}

package() {
  cd "$pkgname"

  python -m installer --destdir="$pkgdir" dist/*.whl

  # license
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
}
