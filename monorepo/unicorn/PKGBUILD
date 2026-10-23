# Maintainer: Levente Polyak <anthraxx[at]archlinux[dot]org>
# Maintainer: Andreas 'Segaja' Schleifer <segaja at archlinux dot org>

pkgbase=unicorn
pkgname=(unicorn python-unicorn ruby-unicorn-engine)
pkgver=2.1.4
pkgrel=2
pkgdesc='Lightweight, multi-platform, multi-architecture CPU emulator framework based on QEMU'
url='https://www.unicorn-engine.org'
arch=(x86_64)
license=(GPL-2.0-only)
makedepends=(
  cmake
  git
  python
  python-build
  python-installer
  python-setuptools
  python-setuptools-scm
  python-versioningit
  python-wheel
  ruby
  ruby-rdoc
)
checkdepends=(
  cmocka
  python-pytest
)
options=('!emptydirs')
source=(
  git+https://github.com/unicorn-engine/unicorn.git#tag=v${pkgver}
)
sha512sums=('4472ae44de5631a3babfd966b0d84910dd512b9e6b3ea8c079eb76b3370c4d13daef5d51741d5cdff6727ee7d1949762999681e83d803df48b37770dff44cd10')
b2sums=('f9b20fa1beb88bd58ea36b417c8f23568cded70952a1cd5bbf7860ca73f3ed1220a038313730c8f48a11a8110c034e2651a15bc870791b35972d8c5618a2caa6')

export SETUPTOOLS_SCM_PRETEND_VERSION=${pkgver}

prepare() {
  cd ${pkgbase}
  sed 's/VERSION =.*/VERSION = "'"${pkgver}"'"/' -i bindings/ruby/unicorn_gem/lib/unicorn_engine/version.rb
}

build() {
  cd ${pkgbase}
  export UNICORN_CFLAGS="${CFLAGS} -ffat-lto-objects"
  export UNICORN_QEMU_FLAGS="--extra-ldflags=\"$LDFLAGS\""
  export QEMU_CXXFLAGS="${CXXFLAGS}"
  export QEMU_LDFLAGS="${LDFLAGS}"
  cmake -B build \
      -DCMAKE_BUILD_TYPE=None \
      -DCMAKE_INSTALL_PREFIX=/usr \
      -DUNICORN_LEGACY_STATIC_ARCHIVE=OFF \
      -Wno-dev
  cmake --build build
  (cd bindings
    python const_generator.py python
    python const_generator.py ruby
  )
  (cd bindings/python
    # avoid rebuilding libunicorn.so
    env LIBUNICORN_PATH="/doesnotexist" python -m build --wheel --no-isolation
    env LIBUNICORN_PATH="/doesnotexist" python -m installer --destdir=tmp_install dist/*.whl
  )
  (cd bindings/ruby/unicorn_gem
    gem build unicorn-engine.gemspec
  )
}

check() {
  cd ${pkgbase}
  ctest --test-dir build --output-on-failure
  cd bindings/python
  local site_packages=$(python -c "import site; print(site.getsitepackages()[0])")
  export LD_LIBRARY_PATH=$PWD/../../build/
  export PYTHONPATH="$PWD/tmp_install/$site_packages"
  py.test -k 'not test_network_auditing and not test_shellcode.py'
  ./tests/test_shellcode.py
  ./tests/test_network_auditing.py
}

package_unicorn() {
  depends=(
    glibc
  )
  provides=(libunicorn.so)
  cd ${pkgbase}
  DESTDIR="${pkgdir}" cmake --install build
  install -Dm 644 samples/*.c -t "${pkgdir}/usr/share/doc/${pkgname}/samples"
}

package_python-unicorn() {
  depends=(
    unicorn
    python
    python-setuptools
  )
  cd ${pkgbase}/bindings/python
  python -m installer --destdir="${pkgdir}" dist/*.whl
}

package_ruby-unicorn-engine() {
  depends=(
    glibc
    unicorn
    ruby
  )
  replaces=('ruby-unicorn')
  cd ${pkgbase}/bindings/ruby/unicorn_gem

  local _gemdir="$(gem env gemdir)"
  gem install \
    --local \
    --verbose \
    --ignore-dependencies \
    --no-user-install \
    --install-dir "${pkgdir}${_gemdir}" \
    --bindir "${pkgdir}/usr/bin" \
    ${pkgbase}-*.gem -- \
    --with-opt-include="${srcdir}/${pkgbase}/include" \
    --with-opt-lib="${srcdir}/${pkgbase}/build"

  install -Dm 644 ../sample* -t "${pkgdir}/usr/share/doc/${pkgname}/samples"

  # remove unrepreducible files
  rm --force --recursive --verbose \
    "${pkgdir}/${_gemdir}/cache/" \
    "${pkgdir}/${_gemdir}/gems/${_gemname}-${pkgver}/vendor/" \
    "${pkgdir}/${_gemdir}/doc/${_gemname}-${pkgver}/ri/ext/"

  find "${pkgdir}/${_gemdir}/gems/" \
    -type f \
    \( \
      -iname "*.o" -o \
      -iname "*.c" -o \
      -iname "*.so" -o \
      -iname "*.time" -o \
      -iname "gem.build_complete" -o \
      -iname "Makefile" \
    \) \
    -delete

  find "${pkgdir}/${_gemdir}/extensions/" \
    -type f \
    \( \
      -iname "mkmf.log" -o \
      -iname "gem_make.out" \
    \) \
    -delete
}

# vim: ts=2 sw=2 et:
