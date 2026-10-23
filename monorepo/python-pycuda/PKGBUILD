# Maintainer: Torsten Keßler <tpkessler@archlinux.org>
# Maintainer: Sven-Hendrik Haase <svenstaro@archlinux.org>
# Contributor: Felix Yan <felixonmars@archlinux.org>
# Contributor: Stéphane Gaudreault <stephane@archlinux.org>

_name=pycuda
pkgname=python-pycuda
pkgver=2026.1
pkgrel=4
pkgdesc="Python wrapper for Nvidia CUDA"
arch=(x86_64)
url="https://documen.tician.de/pycuda/"
license=(
  Apache-2.0
  MIT
)
depends=(
  boost-libs
  cuda
  libgcc
  libstdc++
  glibc
  nvidia-utils
  python
  python-numpy
  python-pytools
)
makedepends=(
  boost
  ctags
  git
  mesa
  python-build
  python-installer
  python-setuptools
  python-wheel
)
checkdepends=(
  python-pytest
)
provides=(pycuda-headers)
conflicts=(pycuda-headers)
replaces=(pycuda-headers)
source=("git+https://github.com/inducer/pycuda.git#tag=v${pkgver}"
        git+https://github.com/inducer/compyte)
sha512sums=('048558cf0f671fe9eec51dcf6282e7017fcd120dad28cb2c414a0f1ac252a726fdb3af63dab0a9b4065cbffb507f898750efd3bbf8bde818dc556ac8c03172e9'
            'SKIP')
b2sums=('ef311ec65f9a1bd4b40072fd18b9abdafea0d1f8c8277ffa7b1a2068c4edb08ebb5d4fef3490e8f45691f08c3f07f2837e434a6bdc44b469bd33c8e545354818'
        'SKIP')

prepare() {
  cd $_name
  git submodule init
  git config submodule.pycuda/compyte.url ../compyte
  git -c protocol.file.allow=always submodule update

  local lib_arch=''
  [[ "$CARCH" = "x86_64" ]] && lib_arch='64'

  # create local siteconf.py for build system
  # defaults can be compared from running the ./configure.py script
  {
    printf "BOOST_PYTHON_LIBNAME = ['boost_python3']\n"
    printf "CUDA_ROOT = '/opt/cuda'\n"
    printf "USE_SHIPPED_BOOST = False\n"
  } > siteconf.py

  # we ship python-numpy
  sed -e 's/oldest-supported-numpy/numpy/' -i pyproject.toml
}

build() {
  cd $_name
  python -m build --wheel --no-isolation
}

# requires an Nvidia graphics card to run
# check() {
#   local pytest_options=(
#     -vv
#     --ignore examples/from-wiki/simple_speed_test.py
#   )
#   local site_packages=$(python -c "import site; print(site.getsitepackages()[0])")

#   cd $_name
#   # install to temporary location, as importlib is used
#   python -m installer --destdir=test_dir dist/*.whl
#   export PYTHONPATH="test_dir/$site_packages:$PYTHONPATH"
#   pytest "${pytest_options[@]}"
# }

package() {
  cd $_name
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm 644 LICENSE -t "$pkgdir"/usr/share/licenses/$pkgname/
}
